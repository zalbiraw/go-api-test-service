package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/notifications/graph/model"
)

// FindNotificationByID is the resolver for the findNotificationByID field.
func (r *entityResolver) FindNotificationByID(ctx context.Context, id string) (*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
