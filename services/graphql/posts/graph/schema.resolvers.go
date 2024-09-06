package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/helpers"

	"go.opentelemetry.io/otel"
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
	tracer := otel.Tracer("post-graphql-api")
	ctx, span := tracer.Start(ctx, "PostResolver")
	defer span.End()

	postId, err := strconv.Atoi(id)
	if nil != err {
		span.RecordError(err)
		return nil, err
	}

	posts := helpers.GetPosts()
	if postId <= 0 || postId > len(posts) {
		err := fmt.Errorf("post not found")
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent(fmt.Sprintf("Fetched post successfully, (%s, %s)", posts[postId].ID, posts[postId].UserID))

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
	tracer := otel.Tracer("posts-graphql-api")
	ctx, span := tracer.Start(ctx, "PostsResolver")
	defer span.End()

	posts := helpers.GetPosts()

	span.AddEvent("Fetched posts successfully")
	return posts, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
