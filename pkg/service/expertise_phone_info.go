package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertisePhoneInfoService struct {
	expertisePhoneInfoStore repositories.ExpertisePhoneInfoStore
}

type ExpertisePhoneInfoService interface {
	Create(model *db.ExpertisePhoneInfo) error
}

func NewExpertisePhoneInfoService(expertisePhoneInfoStore repositories.ExpertisePhoneInfoStore) ExpertisePhoneInfoService {
	return &expertisePhoneInfoService{
		expertisePhoneInfoStore: expertisePhoneInfoStore,
	}
}

func (e *expertisePhoneInfoService) Create(model *db.ExpertisePhoneInfo) error {
	return e.expertisePhoneInfoStore.Create(model)
}
