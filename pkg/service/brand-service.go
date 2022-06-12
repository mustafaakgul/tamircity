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
	brandRepository repositories.BrandRepository
}

func NewBrandService(brandRepository repositories.BrandRepository) BrandService {
	return &brandService{brandRepository: brandRepository}
}
func (m *brandService) Create(model *db.Brand) error {
	return m.brandRepository.Create(model)
}
func (m *brandService) Update(model *db.Brand) error {
	return m.brandRepository.Update(model)
}
func (m *brandService) Delete(model *db.Brand) error {
	return m.brandRepository.Delete(model)
}
func (m *brandService) FindAll() ([]db.Brand, error) {
	return m.brandRepository.FindAll()
}
func (m *brandService) FindByID(id int) (db.Brand, error) {
	return m.brandRepository.FindByID(id)
}
func (m *brandService) FindBy(column string, value interface{}) ([]db.Brand, error) {
	return m.brandRepository.FindBy(column, value)
}
func (m *brandService) Search(query string) ([]db.Brand, error) {
	return m.brandRepository.Search(query)
}
