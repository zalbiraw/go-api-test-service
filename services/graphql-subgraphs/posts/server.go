//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/helpers"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const defaultPort = "4202"

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
			semconv.ServiceNameKey.String("posts-subgraph-service"),
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

	// Initialize and defer tracer shutdown
	shutdown := initTracer()
	defer shutdown()

	// Load posts data
	err := helpers.LoadPosts()
	if err != nil {
		panic("Unable to load posts.")
	}

	// Wrap handlers with OpenTelemetry instrumentation
	http.Handle("/", otelhttp.NewHandler(playground.Handler("GraphQL playground", "/query"), "Playground"))
	http.Handle("/query", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "http-request", r)
		graph.GraphQLEndpointHandler(graph.EndpointOptions{EnableDebug: false}).ServeHTTP(w, r.WithContext(ctx))
	}), "GraphQL Query"))

	// Start the server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
