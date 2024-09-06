package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("posts-subgraph-api")
	ctx, span := tracer.Start(ctx, "PostResolver")
	defer span.End()

	// Convert string ID to integer
	postId, err := strconv.Atoi(id)
	if err != nil {
		span.SetStatus(codes.Error, "Invalid post ID")
		span.RecordError(err)
		return nil, err
	}

	// Fetch posts data
	posts := helpers.GetPosts()

	// Ensure the post ID is valid
	if postId <= 0 || postId > len(posts) {
		err := fmt.Errorf("post not found")
		span.SetStatus(codes.Error, "Post not found")
		span.RecordError(err)
		return nil, err
	}

	// Add trace event on successful fetch
	span.AddEvent(fmt.Sprintf("Fetched post successfully, ID: %d", postId))

	return posts[postId-1], nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("posts-subgraph-api")
	ctx, span := tracer.Start(ctx, "PostsResolver")
	defer span.End()

	// Fetch all posts
	posts := helpers.GetPosts()

	// Add event on successful fetch
	span.AddEvent("Fetched all posts successfully")
	return posts, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
