package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseTvInfoStore struct {
	db *gorm.DB
}

type ExpertiseTvInfoStore interface {
	Create(model *db.ExpertiseTvInfo) error
}

func NewExpertiseTvInfoStore(db *gorm.DB) ExpertiseTvInfoStore {
	return &expertiseTvInfoStore{db: db}
}

func (c *expertiseTvInfoStore) Create(model *db.ExpertiseTvInfo) error {
	return c.db.Create(model).Error
}
