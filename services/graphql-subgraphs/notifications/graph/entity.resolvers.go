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
)

func (r *entityResolver) FindNotificationByID(ctx context.Context, id string) (*model.Notification, error) {
	return &model.Notification{
		ID:     id,
		UserID: strconv.Itoa(rand.Intn(10) + 1),
		Title:  *helpers.RandSentence(),
		Body:   *helpers.RandSentences(5),
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	userId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	var notifications []*model.Notification
	for i := 0; i < userId; i++ {
		notifications = append(notifications, &model.Notification{
			ID:     strconv.Itoa(rand.Intn(1000) + 1),
			UserID: strconv.Itoa(userId),
			Title:  *helpers.RandSentence(),
			Body:   *helpers.RandSentences(5),
		})
	}

	return &model.User{
		ID:            id,
		Notifications: notifications,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
