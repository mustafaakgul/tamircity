package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertiseConsoleInfoService struct {
	expertiseConsoleInfoStore repositories.ExpertiseConsoleInfoStore
}

type ExpertiseConsoleInfoService interface {
	Create(model *db.ExpertiseConsoleInfo) error
	GetByID(id int) (*db.ExpertiseConsoleInfo, error)
}

func NewExpertiseConsoleInfoService(expertiseConsoleInfoStore repositories.ExpertiseConsoleInfoStore) ExpertiseConsoleInfoService {
	return &expertiseConsoleInfoService{
		expertiseConsoleInfoStore: expertiseConsoleInfoStore,
	}
}

func (e *expertiseConsoleInfoService) GetByID(id int) (*db.ExpertiseConsoleInfo, error) {
	return e.expertiseConsoleInfoStore.GetByID(id)
}

func (e *expertiseConsoleInfoService) Create(model *db.ExpertiseConsoleInfo) error {
	return e.expertiseConsoleInfoStore.Create(model)
}
