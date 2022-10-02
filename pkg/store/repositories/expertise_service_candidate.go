package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseServiceCandidateStore struct {
	db *gorm.DB
}

type ExpertiseServiceCandidateStore interface {
	Create(model *db.ExpertiseServiceCandidate) error
	FindByID(id int) (db.ExpertiseServiceCandidate, error)
	Update(model *db.ExpertiseServiceCandidate) error
	//UpdateStatus(model *db.ExpertiseServiceCandidate) error
}

func NewExpertiseServiceCandidateStore(db *gorm.DB) ExpertiseServiceCandidateStore {
	return &expertiseServiceCandidateStore{db: db}
}

func (c *expertiseServiceCandidateStore) Create(model *db.ExpertiseServiceCandidate) error {
	return c.db.Create(model).Error
}

func (c *expertiseServiceCandidateStore) FindByID(id int) (db.ExpertiseServiceCandidate, error) {
	var model db.ExpertiseServiceCandidate
	err := c.db.First(&model, id).Error
	return model, err
}

func (c *expertiseServiceCandidateStore) Update(model *db.ExpertiseServiceCandidate) error {
	return c.db.Save(model).Error
}
