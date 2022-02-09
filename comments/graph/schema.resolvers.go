package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/comments/graph/generated"
	"github.com/zalbiraw/go-api-test-service/comments/graph/helpers"
	"github.com/zalbiraw/go-api-test-service/comments/graph/model"
)

func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	commentId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	comments, err := helpers.GetComments()
	if nil != err {
		return nil, err
	}

	return &((*comments)[commentId]), nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	commentsArray, err := helpers.GetComments()
	if nil != err {
		return nil, err
	}

	comments := make([]*model.Comment, len(*commentsArray))
	for i := 0; i < len(*commentsArray); i++ {
		comments[i] = &((*commentsArray)[i])
	}

	return comments, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
