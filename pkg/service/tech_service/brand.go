package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type BrandService interface {
	Create(model *tech_service.Brand) error
	Update(model *tech_service.Brand) error
	Delete(model *tech_service.Brand) error
	FindAll() ([]tech_service.Brand, error)
	FindByID(id int) (tech_service.Brand, error)
	FindBy(column string, value interface{}) ([]tech_service.Brand, error)
	FindByDeviceTypeId(deviceTypeId int) ([]tech_service3.BrandResponse, error)
	Search(query string) ([]tech_service.Brand, error)
}

type brandService struct {
	brandStore tech_service2.BrandStore
}

func NewBrandService(brandStore tech_service2.BrandStore) BrandService {
	return &brandService{brandStore: brandStore}
}

func (m *brandService) Create(model *tech_service.Brand) error {
	return m.brandStore.Create(model)
}

func (m *brandService) Update(model *tech_service.Brand) error {
	return m.brandStore.Update(model)
}

func (m *brandService) Delete(model *tech_service.Brand) error {
	return m.brandStore.Delete(model)
}

func (m *brandService) FindAll() ([]tech_service.Brand, error) {
	return m.brandStore.FindAll()
}

func (m *brandService) FindByID(id int) (tech_service.Brand, error) {
	return m.brandStore.FindByID(id)
}

func (m *brandService) FindBy(column string, value interface{}) ([]tech_service.Brand, error) {
	return m.brandStore.FindBy(column, value)
}

func (m *brandService) FindByDeviceTypeId(deviceTypeId int) (res []tech_service3.BrandResponse, err error) {
	brands, err := m.brandStore.FindByDeviceTypeId(deviceTypeId)
	if err != nil {
		return nil, err
	}

	for _, brand := range brands {
		res = append(res, tech_service3.BrandResponse{Id: brand.ID, Name: brand.Name})
	}
	return res, nil
}

func (m *brandService) Search(query string) ([]tech_service.Brand, error) {
	return m.brandStore.Search(query)
}
