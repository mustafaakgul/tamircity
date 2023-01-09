package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseWatchInfoStore struct {
	db *gorm.DB
}

type ExpertiseWatchInfoStore interface {
	Create(model *db.ExpertiseWatchInfo) error
}

func NewExpertiseWatchInfoStoreStore(db *gorm.DB) ExpertiseWatchInfoStore {
	return &expertiseWatchInfoStore{db: db}
}

func (c *expertiseWatchInfoStore) Create(model *db.ExpertiseWatchInfo) error {
	return c.db.Create(model).Error
}
