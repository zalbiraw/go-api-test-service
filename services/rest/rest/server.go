package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const defaultPort = "3100"

func initTracer() func() {
	ctx := context.Background()

	// Set up OTLP HTTP exporter
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		log.Fatalf("failed to create OTLP trace exporter: %v", err)
	}

	// Create a new trace provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("json-api-service"),
		)),
	)

	// Register as global trace provider
	otel.SetTracerProvider(tp)

	// Set global propagator for traceparent header support
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	shutdown := initTracer()
	defer shutdown()

	muxer := chi.NewMux()

	muxer.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		tracer := otel.Tracer("json-api-tracer")

		// Start a new span
		ctx, span := tracer.Start(ctx, "HandleJSONRequest")
		defer span.End()

		// Add an attribute to the span
		span.SetAttributes(
			attribute.String("request.path", r.URL.Path),
		)

		// Generate the response
		str := fmt.Sprintf(`{"time":"%s"}`, time.Now().String())
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(str))
		if err != nil {
			span.RecordError(err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}

		span.AddEvent("Successfully handled JSON request")
	})

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
