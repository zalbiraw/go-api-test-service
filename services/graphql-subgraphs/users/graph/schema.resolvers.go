package graph

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/helpers"
)

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("users-subgraph-api")
	ctx, span := tracer.Start(ctx, "UserResolver")
	defer span.End()

	userId, err := strconv.Atoi(id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	users := helpers.GetUsers()
	if userId <= 0 || userId > len(users) {
		err := fmt.Errorf("user not found")
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(
		attribute.Int("user.id", userId),
	)
	span.AddEvent(fmt.Sprintf("Fetched user successfully, (%s, %s)", users[userId-1].Name, users[userId-1].Email))

	return users[userId-1], nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("users-subgraph-api")
	ctx, span := tracer.Start(ctx, "UsersResolver")
	defer span.End()

	users := helpers.GetUsers()

	span.AddEvent("Fetched users successfully")
	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
