package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type technicalServiceCandidateStore struct {
	db *gorm.DB
}

type TechnicalServiceCandidateStore interface {
	Create(model *db.TechnicalServiceCandidate) error
}

func NewTechnicalServiceCandidateStore(db *gorm.DB) TechnicalServiceCandidateStore {
	return &technicalServiceCandidateStore{db: db}
}

func (c *technicalServiceCandidateStore) Create(model *db.TechnicalServiceCandidate) error {
	return c.db.Create(model).Error
}
