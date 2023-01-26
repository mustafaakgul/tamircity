package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertisePcInfoService struct {
	expertisePcInfoStore repositories.ExpertisePcInfoStore
}

type ExpertisePcInfoService interface {
	Create(model *db.ExpertisePcInfo) error
	GetByID(id int) (*db.ExpertisePcInfo, error)
}

func NewExpertisePCInfoService(expertisePcInfoStore repositories.ExpertisePcInfoStore) ExpertisePcInfoService {
	return &expertisePcInfoService{
		expertisePcInfoStore: expertisePcInfoStore,
	}
}

func (e *expertisePcInfoService) GetByID(id int) (*db.ExpertisePcInfo, error) {
	return e.expertisePcInfoStore.GetByID(id)
}

func (e *expertisePcInfoService) Create(model *db.ExpertisePcInfo) error {
	return e.expertisePcInfoStore.Create(model)
}
