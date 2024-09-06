package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/services/graphql/comments/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql/comments/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

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
	tracer := otel.Tracer("comment-graphql-api")
	ctx, span := tracer.Start(ctx, "CommentResolver")
	defer span.End()

	commentId, err := strconv.Atoi(id)
	if nil != err {
		span.RecordError(err)
		return nil, err
	}

	comments := helpers.GetComments()
	if commentId <= 0 || commentId > len(comments) {
		err := fmt.Errorf("comment not found")
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent(fmt.Sprintf("Fetched comment successfully, (%s, %s)", comments[commentId].Name, comments[commentId].Email))

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
	tracer := otel.Tracer("comments-graphql-api")
	ctx, span := tracer.Start(ctx, "CommentsResolver")
	defer span.End()

	comments := helpers.GetComments()

	span.AddEvent("Fetched comments successfully")
	return comments, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
