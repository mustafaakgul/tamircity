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
	FindByID(id int) (db.TechnicalServiceCandidate, error)
	Update(model *db.TechnicalServiceCandidate) error
	//UpdateStatus(model *db.TechnicalServiceCandidate) error
}

func NewTechnicalServiceCandidateStore(db *gorm.DB) TechnicalServiceCandidateStore {
	return &technicalServiceCandidateStore{db: db}
}

func (c *technicalServiceCandidateStore) Create(model *db.TechnicalServiceCandidate) error {
	return c.db.Create(model).Error
}

func (c *technicalServiceCandidateStore) FindByID(id int) (db.TechnicalServiceCandidate, error) {
	var model db.TechnicalServiceCandidate
	err := c.db.First(&model, id).Error
	return model, err
}

func (c *technicalServiceCandidateStore) Update(model *db.TechnicalServiceCandidate) error {
	return c.db.Save(model).Error
}
