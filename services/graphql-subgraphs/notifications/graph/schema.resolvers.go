package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/helpers"
	"math/rand"
	"strconv"
	"time"
)

func (r *subscriptionResolver) GetUserNotifications(ctx context.Context, userID string) (<-chan *model.User, error) {
	userId, err := strconv.Atoi(userID)
	if nil != err {
		return nil, err
	}

	msgs := make(chan *model.User, 1)

	go func() {
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

			msgs <- &model.User{
				ID:            userID,
				Notifications: notifications,
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return msgs, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
