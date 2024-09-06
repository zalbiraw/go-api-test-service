package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/helpers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// getRequestFromContext extracts the HTTP request from context
func getRequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value("http-request").(*http.Request)
	return req, ok
}

// GetUserNotifications is the resolver for the getUserNotifications field.
func (r *subscriptionResolver) GetUserNotifications(ctx context.Context, userID string) (<-chan *model.User, error) {
	// Extract the HTTP request from context
	req, ok := getRequestFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to extract http request from context")
	}

	// Extract tracing context from headers and propagate it
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(req.Header))

	// Start a new span with the propagated context
	tracer := otel.Tracer("notifications-graphql-subgraph")
	ctx, span := tracer.Start(ctx, "GetUserNotificationsResolver")
	defer span.End()

	userId, err := strconv.Atoi(userID)
	if nil != err {
		span.RecordError(err)
		return nil, err
	}

	// Create a channel for notifications
	msgs := make(chan *model.User, 1)

	go func(ctx context.Context) {
		defer close(msgs)
		for {
			var notifications []*model.Notification
			for i := 0; i < rand.Intn(userId)+1; i++ {
				notifications = append(notifications, &model.Notification{
					ID:     strconv.Itoa(rand.Intn(1000) + 1),
					UserID: strconv.Itoa(userId),
					Title:  *helpers.RandSentence(),
					Body:   *helpers.RandSentences(5),
				})
			}

			// Trace the notification event
			_, notificationSpan := tracer.Start(ctx, "SendingNotification")
			notificationSpan.AddEvent("Notifications generated")
			msgs <- &model.User{
				ID:            userID,
				Notifications: notifications,
			}
			notificationSpan.End()

			time.Sleep(1 * time.Second)
		}
	}(ctx)

	return msgs, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
