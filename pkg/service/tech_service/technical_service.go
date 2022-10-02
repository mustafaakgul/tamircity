package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type technicalServiceService struct {
	technicalServiceStore tech_service2.TechnicalServiceStore
}

type TechnicalServiceService interface {
	Create(model *tech_service.TechnicalService) error
	Update(model *tech_service.TechnicalService) error
	Delete(model *tech_service.TechnicalService) error
	FindAll() ([]tech_service.TechnicalService, error)
	FindByID(id int) (tech_service.TechnicalService, error)
	FindBy(column string, value interface{}) ([]tech_service.TechnicalService, error)
	FindByModelId(modelId int) ([]web.TechnicalServiceSearchResponse, error)
	Search(query string) ([]tech_service.TechnicalService, error)
}

func NewTechnicalServiceService(technicalServiceStore tech_service2.TechnicalServiceStore) TechnicalServiceService {
	return &technicalServiceService{
		technicalServiceStore: technicalServiceStore,
	}
}

func (t *technicalServiceService) Create(model *tech_service.TechnicalService) error {
	return t.technicalServiceStore.Create(model)
}

func (t *technicalServiceService) Update(model *tech_service.TechnicalService) error {
	return t.technicalServiceStore.Update(model)
}

func (t *technicalServiceService) Delete(model *tech_service.TechnicalService) error {
	return t.technicalServiceStore.Delete(model)
}

func (t *technicalServiceService) FindAll() ([]tech_service.TechnicalService, error) {
	return t.technicalServiceStore.FindAll()
}

func (t *technicalServiceService) FindByID(id int) (tech_service.TechnicalService, error) {
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
		technicalServiceResponse.Address = technicalService.Address

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

func (t *technicalServiceService) FindBy(column string, value interface{}) ([]tech_service.TechnicalService, error) {
	return t.technicalServiceStore.FindBy(column, value)
}

func (t *technicalServiceService) Search(query string) ([]tech_service.TechnicalService, error) {
	return t.technicalServiceStore.Search(query)
}
