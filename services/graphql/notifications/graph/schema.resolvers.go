package graph

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/zalbiraw/go-api-test-service/services/graphql/notifications/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql/notifications/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

// Placeholder is the resolver for the placeholder field.
func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	str := "Hello World"
	return &str, nil
}

// GetUserNotifications is the resolver for the getUserNotifications field.
func (r *subscriptionResolver) GetUserNotifications(ctx context.Context, userID string) (<-chan []*model.Notification, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("notifications-graphql-api")
	ctx, span := tracer.Start(ctx, "GetUserNotificationsResolver")
	defer span.End()

	userId, err := strconv.Atoi(userID)
	if nil != err {
		span.RecordError(err)
		return nil, err
	}

	msgs := make(chan []*model.Notification, 1)

	go func() {
		var notifications []*model.Notification
		for {
			notifications = append(notifications, &model.Notification{
				ID:     strconv.Itoa(rand.Intn(1000) + 1),
				UserID: strconv.Itoa(userId),
				Title:  *helpers.RandSentence(),
				Body:   *helpers.RandSentences(5),
			})

			msgs <- notifications
			time.Sleep(1 * time.Second)
		}
	}()

	span.SetAttributes(
		attribute.String("user.id", userID),
		attribute.Int("notifications.count", 1), // This can be updated based on actual notifications
	)

	span.AddEvent("Started streaming notifications")
	return msgs, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
