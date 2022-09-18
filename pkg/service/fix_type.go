package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type FixTypeService interface {
	Create(model *db.FixType) error
	Update(model *db.FixType) error
	Delete(model *db.FixType) error
	FindAll() ([]db.FixType, error)
	FindByID(id int) (db.FixType, error)
	FindBy(column string, value interface{}) ([]db.FixType, error)
	FindByDeviceTypeId(deviceTypeId int) ([]web.FixTypeResponse, error)
	Search(query string) ([]db.FixType, error)
}

type fixTypeService struct {
	fixTypeStore repositories.FixTypeStore
}

func NewFixTypeService(fixTypeStore repositories.FixTypeStore) FixTypeService {
	return &fixTypeService{fixTypeStore: fixTypeStore}
}
func (f *fixTypeService) Create(model *db.FixType) error {
	return f.fixTypeStore.Create(model)
}
func (f *fixTypeService) Update(model *db.FixType) error {
	return f.fixTypeStore.Update(model)
}
func (f *fixTypeService) Delete(model *db.FixType) error {
	return f.fixTypeStore.Delete(model)
}
func (f *fixTypeService) FindAll() ([]db.FixType, error) {
	return f.fixTypeStore.FindAll()
}
func (f *fixTypeService) FindByID(id int) (db.FixType, error) {
	return f.fixTypeStore.FindByID(id)
}
func (f *fixTypeService) FindBy(column string, value interface{}) ([]db.FixType, error) {
	return f.fixTypeStore.FindBy(column, value)
}
func (f *fixTypeService) FindByDeviceTypeId(deviceTypeId int) (res []web.FixTypeResponse, err error) {
	fixTypes, err := f.fixTypeStore.FindByDeviceTypeId(deviceTypeId)
	if err != nil {
		return nil, err
	}

	for _, fixType := range fixTypes {
		res = append(res, web.FixTypeResponse{Id: fixType.ID, Description: fixType.Description, Price: fixType.Price})
	}
	return res, nil
}
func (f *fixTypeService) Search(query string) ([]db.FixType, error) {
	return f.fixTypeStore.Search(query)
}
