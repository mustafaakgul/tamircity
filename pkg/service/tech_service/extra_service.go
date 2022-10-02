package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type extraServiceService struct {
	extraServiceStore tech_service2.ExtraServiceStore
}

type ExtraServiceService interface {
	Create(model *tech_service.ExtraService) error
	Update(model *tech_service.ExtraService) error
	Delete(model *tech_service.ExtraService) error
	FindAll() (response []tech_service3.ExtraServiceResponse, err error)
	FindByID(id int) (tech_service.ExtraService, error)
	FindBy(column string, value interface{}) ([]tech_service.ExtraService, error)
	Search(query string) ([]tech_service.ExtraService, error)
}

func NewExtraServiceService(extraServiceStore tech_service2.ExtraServiceStore) ExtraServiceService {
	return &extraServiceService{
		extraServiceStore: extraServiceStore,
	}
}

func (e *extraServiceService) Create(model *tech_service.ExtraService) error {
	return e.extraServiceStore.Create(model)
}

func (e *extraServiceService) Update(model *tech_service.ExtraService) error {
	return e.extraServiceStore.Update(model)
}

func (e *extraServiceService) Delete(model *tech_service.ExtraService) error {
	return e.extraServiceStore.Delete(model)
}

func (e *extraServiceService) FindAll() (response []tech_service3.ExtraServiceResponse, err error) {
	extraServices, err := e.extraServiceStore.FindAll()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	for _, extraService := range extraServices {
		var extraServiceResponse tech_service3.ExtraServiceResponse
		extraServiceResponse.Id = int(extraService.ID)
		extraServiceResponse.Description = extraService.Description
		extraServiceResponse.Price = extraService.Price

		response = append(response, extraServiceResponse)
	}

	return response, nil
}

func (e *extraServiceService) FindByID(id int) (tech_service.ExtraService, error) {
	return e.extraServiceStore.FindByID(id)
}

func (e *extraServiceService) FindBy(column string, value interface{}) ([]tech_service.ExtraService, error) {
	return e.extraServiceStore.FindBy(column, value)
}

func (e *extraServiceService) Search(query string) ([]tech_service.ExtraService, error) {
	return e.extraServiceStore.Search(query)
}
