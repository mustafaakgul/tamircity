package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
	"strconv"
	"strings"
)

type TechnicalServiceShiftService interface {
	FindByTechnicalServiceId(technicalServiceId int) ([]web.TechnicalServiceShiftResponse, error)
	Create(technicalServiceId int, req []web.TechnicalServiceShiftRequest) error
}

type technicalServiceShiftService struct {
	technicalServiceShiftStore tech_service2.TechnicalServiceShiftStore
}

func NewTechnicalServiceShiftService(technicalServiceShiftStore tech_service2.TechnicalServiceShiftStore) TechnicalServiceShiftService {
	return &technicalServiceShiftService{technicalServiceShiftStore: technicalServiceShiftStore}
}

func (t *technicalServiceShiftService) FindByTechnicalServiceId(technicalServiceId int) (res []web.TechnicalServiceShiftResponse, err error) {
	technicalServiceShifts, err := t.technicalServiceShiftStore.FindByTechnicalServiceId(technicalServiceId)
	if err != nil {
		return nil, err
	}
	for _, technicalServiceShift := range technicalServiceShifts {
		res = append(res, web.TechnicalServiceShiftResponse{
			Id:           technicalServiceShift.ID,
			Day:          technicalServiceShift.Day,
			StartOfShift: strings.Join([]string{strconv.Itoa(technicalServiceShift.StartOfShift), "00"}, ":"),
			EndOfShift:   strings.Join([]string{strconv.Itoa(technicalServiceShift.EndOfShift), "00"}, ":"),
			//StartOfShift: technicalServiceShift.StartOfShift,
			//EndOfShift:   technicalServiceShift.EndOfShift,
		})
	}
	return res, nil
}

func (t *technicalServiceShiftService) Create(technicalServiceId int, req []web.TechnicalServiceShiftRequest) error {

	for _, technicalServiceShift := range req {
		var technicalServiceShiftEntity tech_service.TechnicalServiceShift
		technicalServiceShiftEntity.Day = technicalServiceShift.Day
		//technicalServiceShiftEntity.StartOfShift = technicalServiceShift.StartOfShift.Split(":")[0]
		//technicalServiceShiftEntity.EndOfShift = technicalServiceShift.EndOfShift.Split(":")[0]
		technicalServiceShiftEntity.StartOfShift, _ = strconv.Atoi(strings.Split(technicalServiceShift.StartOfShift, ":")[0])
		technicalServiceShiftEntity.EndOfShift, _ = strconv.Atoi(strings.Split(technicalServiceShift.EndOfShift, ":")[0])

		t.technicalServiceShiftStore.CreateOrUpdate(technicalServiceId, technicalServiceShiftEntity)
	}

	return nil
}
