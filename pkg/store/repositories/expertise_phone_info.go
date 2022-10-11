package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertisePhoneInfoStore struct {
	db *gorm.DB
}

type ExpertisePhoneInfoStore interface {
	Create(model *db.ExpertisePhoneInfo) error
}

func NewExpertisePhoneInfoStore(db *gorm.DB) ExpertisePhoneInfoStore {
	return &expertisePhoneInfoStore{db: db}
}

func (c *expertisePhoneInfoStore) Create(model *db.ExpertisePhoneInfo) error {
	return c.db.Create(model).Error
}
