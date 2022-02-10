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

func (r *entityResolver) FindPostByID(ctx context.Context, id string) (*model.Post, error) {
	commentsArray, err := helpers.GetComments()
	if nil != err {
		return nil, err
	}

	var comment model.Comment
	comments := make([]*model.Comment, len(*commentsArray))
	for i := 0; i < len(*commentsArray); i++ {
		comment = (*commentsArray)[i]
		if id == comment.ID {
			comments[i] = &comment
		}
	}

	return &model.Post{
		ID:       id,
		Comments: comments,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
