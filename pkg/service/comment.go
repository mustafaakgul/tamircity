package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type commentService struct {
	commentStore repositories.CommentStore
}

type CommentService interface {
	Create(*web.CommentRequest) error
}

func NewCommentService(commentStore repositories.CommentStore) CommentService {
	return &commentService{
		commentStore: commentStore,
	}
}

func (c *commentService) Create(commentRequest *web.CommentRequest) error {
	var comment db.Comment
	comment.ExpertiseServiceID = commentRequest.ExpertiseServiceID
	comment.CommentOwner = commentRequest.CommentOwner
	comment.Comment = commentRequest.Comment
	comment.CommentDate = commentRequest.CommentDate
	comment.Rate = commentRequest.Rate

	return c.commentStore.Create(&comment)
}
