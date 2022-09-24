package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type TechnicalServiceShiftService interface {
	FindByTechnicalServiceId(technicalServiceId int) ([]web.TechnicalServiceShiftResponse, error)
	Create(technicalServiceId int, req []web.TechnicalServiceShiftRequest) error
}

type technicalServiceShiftService struct {
	technicalServiceShiftStore repositories.TechnicalServiceShiftStore
}

func NewTechnicalServiceShiftService(technicalServiceShiftStore repositories.TechnicalServiceShiftStore) TechnicalServiceShiftService {
	return &technicalServiceShiftService{technicalServiceShiftStore: technicalServiceShiftStore}
}

func (t *technicalServiceShiftService) FindByTechnicalServiceId(technicalServiceId int) (res []web.TechnicalServiceShiftResponse, err error) {
	technicalServiceShifts, err := t.technicalServiceShiftStore.FindByTechnicalServiceId(technicalServiceId)
	if err != nil {
		return nil, err
	}
	for _, technicalServiceShift := range technicalServiceShifts {
		res = append(res, web.TechnicalServiceShiftResponse{Id: technicalServiceShift.ID, Day: technicalServiceShift.Day, StartOfShift: technicalServiceShift.StartOfShift, EndOfShift: technicalServiceShift.EndOfShift})
	}
	return res, nil
}

func (t *technicalServiceShiftService) Create(technicalServiceId int, req []web.TechnicalServiceShiftRequest) error {

	for _, technicalServiceShift := range req {
		var technicalServiceShiftEntity db.TechnicalServiceShift
		technicalServiceShiftEntity.Day = technicalServiceShift.Day
		technicalServiceShiftEntity.StartOfShift = technicalServiceShift.StartOfShift
		technicalServiceShiftEntity.EndOfShift = technicalServiceShift.EndOfShift

		t.technicalServiceShiftStore.CreateOrUpdate(technicalServiceId, technicalServiceShiftEntity)
	}

	return nil
}
