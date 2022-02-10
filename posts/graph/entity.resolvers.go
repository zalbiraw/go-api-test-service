package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zalbiraw/go-api-test-service/posts/graph/generated"
	"github.com/zalbiraw/go-api-test-service/posts/graph/helpers"
	"github.com/zalbiraw/go-api-test-service/posts/graph/model"
	"strconv"
)

func (r *entityResolver) FindPostByID(ctx context.Context, id string) (*model.Post, error) {
	posts, err := helpers.GetPosts()
	if nil != err {
		return nil, err
	}

	postId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	return &((*posts)[postId]), nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
