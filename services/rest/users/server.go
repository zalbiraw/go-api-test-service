package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/zalbiraw/go-api-test-service/services/rest/users/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const defaultPort = "3101"

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
			semconv.ServiceNameKey.String("users-rest-service"),
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

	err := helpers.LoadUsers()
	if nil != err {
		panic("Unable to load users.")
	}

	users := helpers.GetUsers()

	muxer := chi.NewMux()

	muxer.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		tracer := otel.Tracer("users-rest-api")

		// Start a new span
		ctx, span := tracer.Start(ctx, "GetUsers")
		defer span.End()

		span.AddEvent("Fetching all users")

		jsBytes, err := json.Marshal(users)
		if err != nil {
			span.RecordError(err)
			http.Error(w, "Failed to marshal users", http.StatusInternalServerError)
			return
		}

		span.AddEvent("Fetched users successfully")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	muxer.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		tracer := otel.Tracer("user-rest-api")

		// Start a new span
		ctx, span := tracer.Start(ctx, "GetUserByID")
		defer span.End()

		userIDString := chi.URLParam(r, "id")
		userID, err := strconv.Atoi(userIDString)
		if err != nil || userID < 1 || userID > len(users) {
			span.RecordError(err)
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		span.SetAttributes(
			attribute.Int("user.id", userID),
		)

		jsBytes, err := json.Marshal(users[userID-1])
		if err != nil {
			span.RecordError(err)
			http.Error(w, "Failed to marshal user", http.StatusInternalServerError)
			return
		}

		span.AddEvent("Fetched user successfully")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
