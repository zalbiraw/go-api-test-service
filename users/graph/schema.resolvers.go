package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/users/graph/generated"
	"github.com/zalbiraw/go-api-test-service/users/graph/model"
	"github.com/zalbiraw/go-api-test-service/users/helpers"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	users := helpers.GetUsers()

	return &((*users)[userId]), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	usersArray := helpers.GetUsers()

	users := make([]*model.User, len(*usersArray))
	for i := 0; i < len(*usersArray); i++ {
		users[i] = &((*usersArray)[i])
	}

	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
