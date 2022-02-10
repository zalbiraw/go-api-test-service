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

func (r *entityResolver) FindCommentByID(ctx context.Context, id string) (*model.Comment, error) {
	comments, err := helpers.GetComments()
	if nil != err {
		return nil, err
	}

	commentId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	return &((*comments)[commentId]), nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
