package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type ModelService interface {
	Create(model *tech_service.Model) error
	Update(model *tech_service.Model) error
	Delete(model *tech_service.Model) error
	FindAll() ([]tech_service.Model, error)
	FindByID(id int) (tech_service.Model, error)
	FindBy(column string, value interface{}) ([]tech_service.Model, error)
	FindByBrandIdDeviceTypeId(brandId int, deviceTypeId int) ([]web.ModelResponse, error)
	Search(query string) ([]tech_service.Model, error)
}

type modelService struct {
	modelStore tech_service2.ModelStore
}

func NewModelService(modelStore tech_service2.ModelStore) ModelService {
	return &modelService{modelStore: modelStore}
}

func (m *modelService) Create(model *tech_service.Model) error {
	return m.modelStore.Create(model)
}

func (m *modelService) Update(model *tech_service.Model) error {
	return m.modelStore.Update(model)
}

func (m *modelService) Delete(model *tech_service.Model) error {
	return m.modelStore.Delete(model)
}

func (m *modelService) FindAll() ([]tech_service.Model, error) {
	return m.modelStore.FindAll()
}

func (m *modelService) FindByID(id int) (tech_service.Model, error) {
	return m.modelStore.FindByID(id)
}

func (m *modelService) FindBy(column string, value interface{}) ([]tech_service.Model, error) {
	return m.modelStore.FindBy(column, value)
}

func (m *modelService) FindByBrandIdDeviceTypeId(brandId int, deviceTypeId int) (res []web.ModelResponse, err error) {
	//query := fmt.Sprintf("brand_id = %x AND device_type_id = %x ", brandId, deviceTypeId)

	models, err := m.modelStore.Where(brandId, deviceTypeId)
	if err != nil {
		return nil, err
	}
	for _, model := range models {
		res = append(res, web.ModelResponse{Id: model.ID, Name: model.Name})
	}
	return res, nil
}

func (m *modelService) Search(query string) ([]tech_service.Model, error) {
	return m.modelStore.Search(query)
}
