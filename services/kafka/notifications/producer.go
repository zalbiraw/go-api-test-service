package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/linkedin/goavro"

	"github.com/zalbiraw/go-api-test-service/services/kafka/notifications/helpers"
)

const (
	defaultKafkaBrokerURL  = "localhost:9092"
	defaultKafkaTopic      = "test"
	defaultProduceInterval = "60s"
)

// initTracer initializes OpenTelemetry tracing and returns a shutdown function.
func initTracer() func() {
	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		log.Fatalf("failed to create OTLP trace exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("notifications-kafka-producer"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}
}

// getEnvOrDefault retrieves an environment variable or returns the default value.
func getEnvOrDefault(envVar, defaultValue string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}
	return defaultValue
}

// createProducer sets up and returns a Kafka producer.
func createProducer(brokerURL string) (*kafka.Producer, error) {
	tracer := otel.Tracer("kafka-producer")
	_, span := tracer.Start(context.Background(), "createProducer")
	defer span.End()

	config := &kafka.ConfigMap{"bootstrap.servers": brokerURL}
	return kafka.NewProducer(config)
}

// produceMessage encodes a notification and produces it to Kafka.
func produceMessage(producer *kafka.Producer, codec *goavro.Codec, topic string) error {
	tracer := otel.Tracer("kafka-producer")
	_, span := tracer.Start(context.Background(), "produceMessage")
	defer span.End()

	// Get random notification as a map
	notification := helpers.GetRandomNotification()

	// Encode the message into Avro format
	avroBinary, err := codec.BinaryFromNative(nil, notification)
	if err != nil {
		return fmt.Errorf("failed to encode message: %v", err)
	}

	// Create Kafka message and send it
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          avroBinary,
	}
	return producer.Produce(msg, nil)
}

func getCodec() *goavro.Codec {
	schema := `{
		"type": "record",
		"name": "Notification",
		"fields": [
			{
				"name": "notificationType",
				"type": {
					"type": "enum",
					"name": "NotificationType",
					"symbols": ["LIKE", "SHARE", "COMMENT"]
				}
			},
            {
                "name": "msg",
				"type": "string"
			}
        ]
    }`

	// Create Avro codec
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		log.Fatalf("failed to create Avro codec: %v", err)
	}

	return codec
}

// produceMessagesPeriodically sends a message every minute.
func produceMessagesPeriodically(producer *kafka.Producer, codec *goavro.Codec, topic string, produceInterval string) {
	interval, err := time.ParseDuration(produceInterval)
	if err != nil {
		log.Fatalf("invalid produce interval: %v", err)
	}

	for range time.Tick(interval) {
		if err := produceMessage(producer, codec, topic); err != nil {
			log.Printf("failed to produce message: %v", err)
		} else {
			log.Println("Message produced successfully")
		}
		producer.Flush(15 * 1000)
	}
}

func main() {
	shutdown := initTracer()
	defer shutdown()

	// Load users from file
	err := helpers.LoadUsers()
	if err != nil {
		log.Fatalf("failed to load users: %v", err)
	}

	kafkaBrokerURL := getEnvOrDefault("KAFKA_BROKER_URL", defaultKafkaBrokerURL)
	kafkaTopic := getEnvOrDefault("KAFKA_TOPIC", defaultKafkaTopic)
	produceInterval := getEnvOrDefault("PRODUCE_INTERVAL", defaultProduceInterval)

	// Set up Kafka producer
	producer, err := createProducer(kafkaBrokerURL)
	if err != nil {
		log.Fatalf("failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Start producing messages every minute
	produceMessagesPeriodically(producer, getCodec(), kafkaTopic, produceInterval)
}
