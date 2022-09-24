package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type technicalServiceCandidateService struct {
	technicalServiceCandidateStore repositories.TechnicalServiceCandidateStore
}

type TechnicalServiceCandidateService interface {
	Create(model *db.TechnicalServiceCandidate) error
}

func NewTechnicalServiceCandidateService(technicalServiceCandidateStore repositories.TechnicalServiceCandidateStore) TechnicalServiceCandidateService {
	return &technicalServiceCandidateService{
		technicalServiceCandidateStore: technicalServiceCandidateStore,
	}
}

func (c *technicalServiceCandidateService) Create(model *db.TechnicalServiceCandidate) error {
	return c.technicalServiceCandidateStore.Create(model)
}
