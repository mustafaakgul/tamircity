package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type ModelService interface {
	Create(model *db.Model) error
	Update(model *db.Model) error
	Delete(model *db.Model) error
	FindAll() ([]db.Model, error)
	FindByID(id int) (db.Model, error)
	FindBy(column string, value interface{}) ([]db.Model, error)
	FindByBrandIdDeviceTypeId(brandId int, deviceTypeId int) ([]web.ModelResponse, error)
	Search(query string) ([]db.Model, error)
}

type modelService struct {
	modelStore repositories.ModelStore
}

func NewModelService(modelStore repositories.ModelStore) ModelService {
	return &modelService{modelStore: modelStore}
}

func (m *modelService) Create(model *db.Model) error {
	return m.modelStore.Create(model)
}

func (m *modelService) Update(model *db.Model) error {
	return m.modelStore.Update(model)
}

func (m *modelService) Delete(model *db.Model) error {
	return m.modelStore.Delete(model)
}

func (m *modelService) FindAll() ([]db.Model, error) {
	return m.modelStore.FindAll()
}

func (m *modelService) FindByID(id int) (db.Model, error) {
	return m.modelStore.FindByID(id)
}

func (m *modelService) FindBy(column string, value interface{}) ([]db.Model, error) {
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

func (m *modelService) Search(query string) ([]db.Model, error) {
	return m.modelStore.Search(query)
}
