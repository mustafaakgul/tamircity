package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type expertiseServiceService struct {
	expertiseServiceStore repositories.ExpertiseServiceStore
}

type ExpertiseServiceService interface {
	Create(model *db.ExpertiseService) error
	Update(model *db.ExpertiseService) error
	Delete(model *db.ExpertiseService) error
	FindAll() ([]db.ExpertiseService, error)
	FindByID(id int) (db.ExpertiseService, error)
	FindBy(column string, value interface{}) ([]db.ExpertiseService, error)
	FindByModelId(modelId int) ([]web.ExpertiseServiceSearchResponse, error)
	Search(query string) ([]db.ExpertiseService, error)
}

func NewExpertiseServiceService(expertiseServiceStore repositories.ExpertiseServiceStore) ExpertiseServiceService {
	return &expertiseServiceService{
		expertiseServiceStore: expertiseServiceStore,
	}
}

func (t *expertiseServiceService) Create(model *db.ExpertiseService) error {
	return t.expertiseServiceStore.Create(model)
}

func (t *expertiseServiceService) Update(model *db.ExpertiseService) error {
	return t.expertiseServiceStore.Update(model)
}

func (t *expertiseServiceService) Delete(model *db.ExpertiseService) error {
	return t.expertiseServiceStore.Delete(model)
}

func (t *expertiseServiceService) FindAll() ([]db.ExpertiseService, error) {
	return t.expertiseServiceStore.FindAll()
}

func (t *expertiseServiceService) FindByID(id int) (db.ExpertiseService, error) {
	return t.expertiseServiceStore.FindByID(id)
}

func (t *expertiseServiceService) FindByModelId(modelId int) (response []web.ExpertiseServiceSearchResponse, err error) {
	expertiseServices, err := t.expertiseServiceStore.FindByModelId(modelId)

	if err != nil {
		return nil, err
	}

	for _, expertiseService := range expertiseServices {
		var expertiseServiceResponse web.ExpertiseServiceSearchResponse
		expertiseServiceResponse.Id = int(expertiseService.ID)
		expertiseServiceResponse.Name = expertiseService.ServiceName
		expertiseServiceResponse.Address = expertiseService.Address

		for _, shift := range expertiseService.ExpertiseServiceShifts {
			expertiseServiceResponse.ExpertiseServiceShift.Day = shift.Day
			expertiseServiceResponse.ExpertiseServiceShift.StartOfShift = shift.StartOfShift
			expertiseServiceResponse.ExpertiseServiceShift.EndOfShift = shift.EndOfShift
		}

		// TODO: Error
		/*for _, reservation := range expertiseService.ExpertiseServiceReservations {
			var expertiseServiceReservation web.ExpertiseServiceReservation
			expertiseServiceReservation.Day = reservation.Day
			expertiseServiceReservation.DateOfDay = reservation.DateofDay
			expertiseServiceReservation.StartOfShift = reservation.StartOfShift
			expertiseServiceReservation.EndOfShift = reservation.EndOfShift
			expertiseServiceResponse.ExpertiseServiceReservations = append(expertiseServiceResponse.ExpertiseServiceReservations, expertiseServiceReservation)
		}*/
		response = append(response, expertiseServiceResponse)
	}

	return response, nil
}

func (t *expertiseServiceService) FindBy(column string, value interface{}) ([]db.ExpertiseService, error) {
	return t.expertiseServiceStore.FindBy(column, value)
}

func (t *expertiseServiceService) Search(query string) ([]db.ExpertiseService, error) {
	return t.expertiseServiceStore.Search(query)
}
