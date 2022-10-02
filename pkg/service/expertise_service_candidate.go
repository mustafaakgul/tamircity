package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertiseServiceCandidateService struct {
	expertiseServiceCandidateStore repositories.ExpertiseServiceCandidateStore
}

type ExpertiseServiceCandidateService interface {
	Create(model *db.ExpertiseServiceCandidate) error
	Update(model *db.ExpertiseServiceCandidate) error
	FindByID(id int) (db.ExpertiseServiceCandidate, error)
}

func NewExpertiseServiceCandidateService(expertiseServiceCandidateStore repositories.ExpertiseServiceCandidateStore) ExpertiseServiceCandidateService {
	return &expertiseServiceCandidateService{
		expertiseServiceCandidateStore: expertiseServiceCandidateStore,
	}
}

func (c *expertiseServiceCandidateService) Create(model *db.ExpertiseServiceCandidate) error {
	return c.expertiseServiceCandidateStore.Create(model)
}

func (c *expertiseServiceCandidateService) Update(model *db.ExpertiseServiceCandidate) error {
	return c.expertiseServiceCandidateStore.Update(model)
}

func (c *expertiseServiceCandidateService) FindByID(id int) (db.ExpertiseServiceCandidate, error) {
	return c.expertiseServiceCandidateStore.FindByID(id)
}
