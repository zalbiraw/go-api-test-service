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

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	users := helpers.GetUsers()

	return users[userId-1], nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return helpers.GetUsers(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
