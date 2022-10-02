package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
	"strconv"
	"strings"
)

type ExpertiseServiceShiftService interface {
	FindByExpertiseServiceId(expertiseServiceId int) ([]web.ExpertiseServiceShiftResponse, error)
	Create(expertiseServiceId int, req []web.ExpertiseServiceShiftRequest) error
}

type expertiseServiceShiftService struct {
	expertiseServiceShiftStore repositories.ExpertiseServiceShiftStore
}

func NewExpertiseServiceShiftService(expertiseServiceShiftStore repositories.ExpertiseServiceShiftStore) ExpertiseServiceShiftService {
	return &expertiseServiceShiftService{expertiseServiceShiftStore: expertiseServiceShiftStore}
}

func (t *expertiseServiceShiftService) FindByExpertiseServiceId(expertiseServiceId int) (res []web.ExpertiseServiceShiftResponse, err error) {
	expertiseServiceShifts, err := t.expertiseServiceShiftStore.FindByExpertiseServiceId(expertiseServiceId)
	if err != nil {
		return nil, err
	}
	for _, expertiseServiceShift := range expertiseServiceShifts {
		res = append(res, web.ExpertiseServiceShiftResponse{
			Id:           expertiseServiceShift.ID,
			Day:          expertiseServiceShift.Day,
			StartOfShift: strings.Join([]string{strconv.Itoa(expertiseServiceShift.StartOfShift), "00"}, ":"),
			EndOfShift:   strings.Join([]string{strconv.Itoa(expertiseServiceShift.EndOfShift), "00"}, ":"),
			//StartOfShift: expertiseServiceShift.StartOfShift,
			//EndOfShift:   expertiseServiceShift.EndOfShift,
		})
	}
	return res, nil
}

func (t *expertiseServiceShiftService) Create(expertiseServiceId int, req []web.ExpertiseServiceShiftRequest) error {

	for _, expertiseServiceShift := range req {
		var expertiseServiceShiftEntity db.ExpertiseServiceShift
		expertiseServiceShiftEntity.Day = expertiseServiceShift.Day
		//expertiseServiceShiftEntity.StartOfShift = expertiseServiceShift.StartOfShift.Split(":")[0]
		//expertiseServiceShiftEntity.EndOfShift = expertiseServiceShift.EndOfShift.Split(":")[0]
		expertiseServiceShiftEntity.StartOfShift, _ = strconv.Atoi(strings.Split(expertiseServiceShift.StartOfShift, ":")[0])
		expertiseServiceShiftEntity.EndOfShift, _ = strconv.Atoi(strings.Split(expertiseServiceShift.EndOfShift, ":")[0])

		t.expertiseServiceShiftStore.CreateOrUpdate(expertiseServiceId, expertiseServiceShiftEntity)
	}

	return nil
}
