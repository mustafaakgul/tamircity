package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertiseWatchInfoService struct {
	expertiseWatchInfoStore repositories.ExpertiseWatchInfoStore
}

type ExpertiseWatchInfoService interface {
	Create(model *db.ExpertiseWatchInfo) error
	GetByID(id int) (*db.ExpertiseWatchInfo, error)
}

func NewExpertiseWatchInfoService(expertiseWatchInfoStore repositories.ExpertiseWatchInfoStore) ExpertiseWatchInfoService {
	return &expertiseWatchInfoService{
		expertiseWatchInfoStore: expertiseWatchInfoStore,
	}
}

func (e *expertiseWatchInfoService) GetByID(id int) (*db.ExpertiseWatchInfo, error) {
	return e.expertiseWatchInfoStore.GetByID(id)
}

func (e *expertiseWatchInfoService) Create(model *db.ExpertiseWatchInfo) error {
	return e.expertiseWatchInfoStore.Create(model)
}
