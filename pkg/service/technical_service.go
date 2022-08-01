package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/models/web"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type technicalServiceService struct {
	technicalServiceStore repositories.TechnicalServiceStore
}

type TechnicalServiceService interface {
	Create(model *db.TechnicalService) error
	Update(model *db.TechnicalService) error
	Delete(model *db.TechnicalService) error
	FindAll() ([]db.TechnicalService, error)
	FindByID(id int) (db.TechnicalService, error)
	FindBy(column string, value interface{}) ([]db.TechnicalService, error)
	FindByModelId(modelId int) ([]web.TechnicalServiceSearchResponse, error)
	Search(query string) ([]db.TechnicalService, error)
}

func NewTechnicalServiceService(technicalServiceStore repositories.TechnicalServiceStore) TechnicalServiceService {
	return &technicalServiceService{
		technicalServiceStore: technicalServiceStore,
	}
}

func (t *technicalServiceService) Create(model *db.TechnicalService) error {
	return t.technicalServiceStore.Create(model)
}

func (t *technicalServiceService) Update(model *db.TechnicalService) error {
	return t.technicalServiceStore.Update(model)
}

func (t *technicalServiceService) Delete(model *db.TechnicalService) error {
	return t.technicalServiceStore.Delete(model)
}

func (t *technicalServiceService) FindAll() ([]db.TechnicalService, error) {
	return t.technicalServiceStore.FindAll()
}

func (t *technicalServiceService) FindByID(id int) (db.TechnicalService, error) {
	return t.technicalServiceStore.FindByID(id)
}

func (t *technicalServiceService) FindByModelId(modelId int) (response []web.TechnicalServiceSearchResponse, err error) {
	technicalServices, err := t.technicalServiceStore.FindByModelId(modelId)

	if err != nil {
		return nil, err
	}

	for _, technicalService := range technicalServices {
		var technicalServiceResponse web.TechnicalServiceSearchResponse
		technicalServiceResponse.Id = int(technicalService.ID)
		technicalServiceResponse.Name = technicalService.ServiceName
		technicalServiceResponse.Address = "" // TO DO

		for _, shift := range technicalService.TechnicalServiceShifts {
			technicalServiceResponse.TechnicalServiceShift.Day = shift.Day
			technicalServiceResponse.TechnicalServiceShift.StartOfShift = shift.StartOfShift
			technicalServiceResponse.TechnicalServiceShift.EndOfShift = shift.EndOfShift
		}

		for _, reservation := range technicalService.TechnicalServiceReservations {
			var technicalServiceReservation web.TechnicalServiceReservation
			technicalServiceReservation.Day = reservation.Day
			technicalServiceReservation.DateOfDay = reservation.DateofDay
			technicalServiceReservation.StartOfShift = reservation.StartOfShift
			technicalServiceReservation.EndOfShift = reservation.EndOfShift
			technicalServiceResponse.TechnicalServiceReservations = append(technicalServiceResponse.TechnicalServiceReservations, technicalServiceReservation)
		}
		response = append(response, technicalServiceResponse)
	}

	return response, nil
}

func (t *technicalServiceService) FindBy(column string, value interface{}) ([]db.TechnicalService, error) {
	return t.technicalServiceStore.FindBy(column, value)
}

func (t *technicalServiceService) Search(query string) ([]db.TechnicalService, error) {
	return t.technicalServiceStore.Search(query)
}
