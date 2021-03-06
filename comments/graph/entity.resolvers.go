package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/zalbiraw/go-api-test-service/comments/graph/generated"
	"github.com/zalbiraw/go-api-test-service/comments/graph/model"
	"github.com/zalbiraw/go-api-test-service/comments/helpers"
)

func (r *entityResolver) FindCommentByID(ctx context.Context, id string) (*model.Comment, error) {
	commentId, err := strconv.Atoi(id)
	if nil != err {
		return nil, err
	}

	comments := helpers.GetComments()

	return comments[commentId-1], nil
}

func (r *entityResolver) FindPostByID(ctx context.Context, id string) (*model.Post, error) {
	commentsArray := helpers.GetComments()

	var comments []*model.Comment
	for i := 0; i < len(commentsArray); i++ {
		if id == commentsArray[i].PostID {
			comments = append(comments, commentsArray[i])
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
