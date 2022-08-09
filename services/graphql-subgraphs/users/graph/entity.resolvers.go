package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/model"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	userId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	users := helpers.GetUsers()

	return users[userId-1], nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
