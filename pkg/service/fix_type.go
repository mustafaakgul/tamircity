package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type FixTypeService interface {
	Create(model *db.FixType) error
	Update(model *db.FixType) error
	Delete(model *db.FixType) error
	FindAll() ([]db.FixType, error)
	FindByID(id int) (db.FixType, error)
	FindBy(column string, value interface{}) ([]db.FixType, error)
	Search(query string) ([]db.FixType, error)
}

type fixTypeService struct {
	fixTypeStore repositories.FixTypeStore
}

func NewFixTypeService(fixTypeStore repositories.FixTypeStore) FixTypeService {
	return &fixTypeService{fixTypeStore: fixTypeStore}
}
func (m *fixTypeService) Create(model *db.FixType) error {
	return m.fixTypeStore.Create(model)
}
func (m *fixTypeService) Update(model *db.FixType) error {
	return m.fixTypeStore.Update(model)
}
func (m *fixTypeService) Delete(model *db.FixType) error {
	return m.fixTypeStore.Delete(model)
}
func (m *fixTypeService) FindAll() ([]db.FixType, error) {
	return m.fixTypeStore.FindAll()
}
func (m *fixTypeService) FindByID(id int) (db.FixType, error) {
	return m.fixTypeStore.FindByID(id)
}
func (m *fixTypeService) FindBy(column string, value interface{}) ([]db.FixType, error) {
	return m.fixTypeStore.FindBy(column, value)
}
func (m *fixTypeService) Search(query string) ([]db.FixType, error) {
	return m.fixTypeStore.Search(query)
}
