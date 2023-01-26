package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertiseTvInfoService struct {
	expertiseTvInfoStore repositories.ExpertiseTvInfoStore
}

type ExpertiseTvInfoService interface {
	Create(model *db.ExpertiseTvInfo) error
	GetByID(id int) (*db.ExpertiseTvInfo, error)
}

func NewExpertiseTvInfoService(expertiseTvInfoStore repositories.ExpertiseTvInfoStore) ExpertiseTvInfoService {
	return &expertiseTvInfoService{
		expertiseTvInfoStore: expertiseTvInfoStore,
	}
}

func (e *expertiseTvInfoService) GetByID(id int) (*db.ExpertiseTvInfo, error) {
	return e.expertiseTvInfoStore.GetByID(id)
}

func (e *expertiseTvInfoService) Create(model *db.ExpertiseTvInfo) error {
	return e.expertiseTvInfoStore.Create(model)
}
