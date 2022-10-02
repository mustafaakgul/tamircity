package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type technicalServiceCandidateService struct {
	technicalServiceCandidateStore tech_service2.TechnicalServiceCandidateStore
}

type TechnicalServiceCandidateService interface {
	Create(model *tech_service.TechnicalServiceCandidate) error
	Update(model *tech_service.TechnicalServiceCandidate) error
	FindByID(id int) (tech_service.TechnicalServiceCandidate, error)
}

func NewTechnicalServiceCandidateService(technicalServiceCandidateStore tech_service2.TechnicalServiceCandidateStore) TechnicalServiceCandidateService {
	return &technicalServiceCandidateService{
		technicalServiceCandidateStore: technicalServiceCandidateStore,
	}
}

func (c *technicalServiceCandidateService) Create(model *tech_service.TechnicalServiceCandidate) error {
	return c.technicalServiceCandidateStore.Create(model)
}

func (c *technicalServiceCandidateService) Update(model *tech_service.TechnicalServiceCandidate) error {
	return c.technicalServiceCandidateStore.Update(model)
}

func (c *technicalServiceCandidateService) FindByID(id int) (tech_service.TechnicalServiceCandidate, error) {
	return c.technicalServiceCandidateStore.FindByID(id)
}
