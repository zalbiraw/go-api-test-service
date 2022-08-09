package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zalbiraw/go-api-test-service/helpers"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/generated"
	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/posts/graph/model"
	"strconv"
)

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	postId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	posts := helpers.GetPosts()

	return posts[postId-1], nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return helpers.GetPosts(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
