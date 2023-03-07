package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type commentStore struct {
	db *gorm.DB
}

type CommentStore interface {
	Create(comment *db.Comment) error
}

func NewCommentStore(db *gorm.DB) CommentStore {
	return &commentStore{db: db}
}

func (c *commentStore) Create(comment *db.Comment) error {
	return c.db.Create(comment).Error
}
