package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type expertiseServiceCandidateStore struct {
	db *gorm.DB
}

type ExpertiseServiceCandidateStore interface {
	Create(model *expr_service.ExpertiseServiceCandidate) error
	FindByID(id int) (expr_service.ExpertiseServiceCandidate, error)
	Update(model *expr_service.ExpertiseServiceCandidate) error
	//UpdateStatus(model *db.ExpertiseServiceCandidate) error
}

func NewExpertiseServiceCandidateStore(db *gorm.DB) ExpertiseServiceCandidateStore {
	return &expertiseServiceCandidateStore{db: db}
}

func (c *expertiseServiceCandidateStore) Create(model *expr_service.ExpertiseServiceCandidate) error {
	return c.db.Create(model).Error
}

func (c *expertiseServiceCandidateStore) FindByID(id int) (expr_service.ExpertiseServiceCandidate, error) {
	var model expr_service.ExpertiseServiceCandidate
	err := c.db.First(&model, id).Error
	return model, err
}

func (c *expertiseServiceCandidateStore) Update(model *expr_service.ExpertiseServiceCandidate) error {
	return c.db.Save(model).Error
}
