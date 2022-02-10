package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zalbiraw/go-api-test-service/users/graph/generated"
	"github.com/zalbiraw/go-api-test-service/users/graph/helpers"
	"github.com/zalbiraw/go-api-test-service/users/graph/model"
	"strconv"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	users, err := helpers.GetUsers()
	if nil != err {
		return nil, err
	}

	userId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	return &((*users)[userId]), nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
