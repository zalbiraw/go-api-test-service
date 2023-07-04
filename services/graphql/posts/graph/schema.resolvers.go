package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/graph/model"
	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/helpers"
)

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	postId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	posts := helpers.GetPosts()

	return posts[postId-1], nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return helpers.GetPosts(), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
