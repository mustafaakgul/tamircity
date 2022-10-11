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
}

func NewExpertiseTvInfoService(expertiseTvInfoStore repositories.ExpertiseTvInfoStore) ExpertiseTvInfoService {
	return &expertiseTvInfoService{
		expertiseTvInfoStore: expertiseTvInfoStore,
	}
}

func (e *expertiseTvInfoService) Create(model *db.ExpertiseTvInfo) error {
	return e.expertiseTvInfoStore.Create(model)
}
