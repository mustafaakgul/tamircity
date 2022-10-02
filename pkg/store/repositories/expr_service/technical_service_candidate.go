package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"gorm.io/gorm"
)

type technicalServiceCandidateStore struct {
	db *gorm.DB
}

type TechnicalServiceCandidateStore interface {
	Create(model *tech_service.TechnicalServiceCandidate) error
	FindByID(id int) (tech_service.TechnicalServiceCandidate, error)
	Update(model *tech_service.TechnicalServiceCandidate) error
	//UpdateStatus(model *db.TechnicalServiceCandidate) error
}

func NewTechnicalServiceCandidateStore(db *gorm.DB) TechnicalServiceCandidateStore {
	return &technicalServiceCandidateStore{db: db}
}

func (c *technicalServiceCandidateStore) Create(model *tech_service.TechnicalServiceCandidate) error {
	return c.db.Create(model).Error
}

func (c *technicalServiceCandidateStore) FindByID(id int) (tech_service.TechnicalServiceCandidate, error) {
	var model tech_service.TechnicalServiceCandidate
	err := c.db.First(&model, id).Error
	return model, err
}

func (c *technicalServiceCandidateStore) Update(model *tech_service.TechnicalServiceCandidate) error {
	return c.db.Save(model).Error
}
