package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/comments/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("comments-graphql-subgraph")
	ctx, span := tracer.Start(ctx, "CommentQuery")
	defer span.End()

	// Convert ID from string to integer
	commentId, err := strconv.Atoi(id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	// Retrieve comments
	comments := helpers.GetComments()

	// Check bounds and return the comment
	if commentId <= 0 || commentId > len(comments) {
		err := errors.New("comment ID out of range")
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent("Successfully retrieved comment")
	return comments[commentId-1], nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("comments-graphql-subgraph")
	ctx, span := tracer.Start(ctx, "CommentsQuery")
	defer span.End()

	// Retrieve comments
	comments := helpers.GetComments()

	span.AddEvent("Successfully retrieved comments list")
	return comments, nil
}

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
