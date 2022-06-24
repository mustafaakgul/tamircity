package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type BrandService interface {
	Create(model *db.Brand) error
	Update(model *db.Brand) error
	Delete(model *db.Brand) error
	FindAll() ([]db.Brand, error)
	FindByID(id int) (db.Brand, error)
	FindBy(column string, value interface{}) ([]db.Brand, error)
	Search(query string) ([]db.Brand, error)
}

type brandService struct {
	brandStore repositories.BrandStore
}

func NewBrandService(brandStore repositories.BrandStore) BrandService {
	return &brandService{brandStore: brandStore}
}
func (m *brandService) Create(model *db.Brand) error {
	return m.brandStore.Create(model)
}
func (m *brandService) Update(model *db.Brand) error {
	return m.brandStore.Update(model)
}
func (m *brandService) Delete(model *db.Brand) error {
	return m.brandStore.Delete(model)
}
func (m *brandService) FindAll() ([]db.Brand, error) {
	return m.brandStore.FindAll()
}
func (m *brandService) FindByID(id int) (db.Brand, error) {
	return m.brandStore.FindByID(id)
}
func (m *brandService) FindBy(column string, value interface{}) ([]db.Brand, error) {
	return m.brandStore.FindBy(column, value)
}
func (m *brandService) Search(query string) ([]db.Brand, error) {
	return m.brandStore.Search(query)
}
