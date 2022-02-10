package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/posts/graph/generated"
	"github.com/zalbiraw/go-api-test-service/posts/graph/helpers"
	"github.com/zalbiraw/go-api-test-service/posts/graph/model"
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

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	postsArray, err := helpers.GetPosts()
	if nil != err {
		return nil, err
	}

	var post model.Post
	posts := make([]*model.Post, len(*postsArray))
	for i := 0; i < len(*postsArray); i++ {
		post = (*postsArray)[i]
		if id == post.ID {
			posts[i] = &post
		}
	}

	return &model.User{
		ID:    id,
		Posts: posts,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
