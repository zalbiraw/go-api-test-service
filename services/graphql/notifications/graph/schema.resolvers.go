package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/zalbiraw/go-api-test-service/services/graphql/notifications/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql/notifications/helpers"
)

// GetUserNotifications is the resolver for the getUserNotifications field.
func (r *subscriptionResolver) GetUserNotifications(ctx context.Context, userID string) (<-chan []*model.Notification, error) {
	userId, err := strconv.Atoi(userID)
	if nil != err {
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

	return msgs, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
