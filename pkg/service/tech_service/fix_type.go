package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type FixTypeService interface {
	Create(model *tech_service.FixType) error
	Update(model *tech_service.FixType) error
	Delete(model *tech_service.FixType) error
	FindAll() ([]tech_service.FixType, error)
	FindByID(id int) (tech_service.FixType, error)
	FindBy(column string, value interface{}) ([]tech_service.FixType, error)
	FindByDeviceTypeId(deviceTypeId int) ([]tech_service3.FixTypeResponse, error)
	Search(query string) ([]tech_service.FixType, error)
}

type fixTypeService struct {
	fixTypeStore tech_service2.FixTypeStore
}

func NewFixTypeService(fixTypeStore tech_service2.FixTypeStore) FixTypeService {
	return &fixTypeService{fixTypeStore: fixTypeStore}
}
func (f *fixTypeService) Create(model *tech_service.FixType) error {
	return f.fixTypeStore.Create(model)
}
func (f *fixTypeService) Update(model *tech_service.FixType) error {
	return f.fixTypeStore.Update(model)
}
func (f *fixTypeService) Delete(model *tech_service.FixType) error {
	return f.fixTypeStore.Delete(model)
}
func (f *fixTypeService) FindAll() ([]tech_service.FixType, error) {
	return f.fixTypeStore.FindAll()
}
func (f *fixTypeService) FindByID(id int) (tech_service.FixType, error) {
	return f.fixTypeStore.FindByID(id)
}
func (f *fixTypeService) FindBy(column string, value interface{}) ([]tech_service.FixType, error) {
	return f.fixTypeStore.FindBy(column, value)
}
func (f *fixTypeService) FindByDeviceTypeId(deviceTypeId int) (res []tech_service3.FixTypeResponse, err error) {
	fixTypes, err := f.fixTypeStore.FindByDeviceTypeId(deviceTypeId)
	if err != nil {
		return nil, err
	}

	for _, fixType := range fixTypes {
		res = append(res, tech_service3.FixTypeResponse{Id: fixType.ID, Description: fixType.Description, Price: fixType.Price})
	}
	return res, nil
}
func (f *fixTypeService) Search(query string) ([]tech_service.FixType, error) {
	return f.fixTypeStore.Search(query)
}
