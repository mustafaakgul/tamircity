package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertisePcInfoStore struct {
	db *gorm.DB
}

type ExpertisePcInfoStore interface {
	Create(model *db.ExpertisePcInfo) error
}

func NewExpertisePcInfoStoreStore(db *gorm.DB) ExpertisePcInfoStore {
	return &expertisePcInfoStore{db: db}
}

func (c *expertisePcInfoStore) Create(model *db.ExpertisePcInfo) error {
	return c.db.Create(model).Error
}
