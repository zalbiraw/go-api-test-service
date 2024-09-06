package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/zalbiraw/go-api-test-service/services/rest/comments/helpers"
	"github.com/zalbiraw/go-api-test-service/services/rest/comments/model"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const defaultPort = "3103"

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
			semconv.ServiceNameKey.String("comments-rest-service"),
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

	err := helpers.LoadComments()
	if nil != err {
		panic("Unable to load comments.")
	}

	comments := helpers.GetComments()

	muxer := chi.NewMux()

	muxer.Get("/posts/{id}/comments", func(w http.ResponseWriter, r *http.Request) {
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		tracer := otel.Tracer("comments-rest-api")

		// Start a new span
		ctx, span := tracer.Start(ctx, "GetCommentsByPostID")
		defer span.End()

		postID := chi.URLParam(r, "id")

		var postComments []*model.Comment

		for _, comment := range comments {
			if comment.PostID == postID {
				postComments = append(postComments, comment)
			}
		}

		span.SetAttributes(
			attribute.String("post.id", postID),
			attribute.Int("comment.count", len(postComments)),
		)

		jsBytes, err := json.Marshal(postComments)
		if err != nil {
			span.RecordError(err)
			http.Error(w, "Failed to marshal comments", http.StatusInternalServerError)
			return
		}

		span.AddEvent("Fetched comments successfully")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsBytes)
	})

	log.Printf("connect to http://localhost:%s/ to test API", port)
	log.Fatal(http.ListenAndServe(":"+port, muxer))
}
