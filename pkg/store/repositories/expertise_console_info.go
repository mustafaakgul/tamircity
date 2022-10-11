package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseConsoleInfoStore struct {
	db *gorm.DB
}

type ExpertiseConsoleInfoStore interface {
	Create(model *db.ExpertiseConsoleInfo) error
}

func NewExpertiseConsoleInfoStore(db *gorm.DB) ExpertiseConsoleInfoStore {
	return &expertiseConsoleInfoStore{db: db}
}

func (c *expertiseConsoleInfoStore) Create(model *db.ExpertiseConsoleInfo) error {
	return c.db.Create(model).Error
}
