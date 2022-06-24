package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type technicalServiceService struct {
	technicalServiceStore repositories.TechnicalServiceStore
}

type TechnicalServiceService interface {
	Create(model *db.TechnicalService) error
	Update(model *db.TechnicalService) error
	Delete(model *db.TechnicalService) error
	FindAll() ([]db.TechnicalService, error)
	FindByID(id int) (db.TechnicalService, error)
	FindBy(column string, value interface{}) ([]db.TechnicalService, error)
	Search(query string) ([]db.TechnicalService, error)
}

func NewTechnicalServiceService(technicalServiceStore repositories.TechnicalServiceStore) TechnicalServiceService {
	return &technicalServiceService{
		technicalServiceStore: technicalServiceStore,
	}
}

func (t *technicalServiceService) Create(model *db.TechnicalService) error {
	return t.technicalServiceStore.Create(model)
}

func (t *technicalServiceService) Update(model *db.TechnicalService) error {
	return t.technicalServiceStore.Update(model)
}

func (t *technicalServiceService) Delete(model *db.TechnicalService) error {
	return t.technicalServiceStore.Delete(model)
}

func (t *technicalServiceService) FindAll() ([]db.TechnicalService, error) {
	return t.technicalServiceStore.FindAll()
}

func (t *technicalServiceService) FindByID(id int) (db.TechnicalService, error) {
	return t.technicalServiceStore.FindByID(id)
}

func (t *technicalServiceService) FindBy(column string, value interface{}) ([]db.TechnicalService, error) {
	return t.technicalServiceStore.FindBy(column, value)
}

func (t *technicalServiceService) Search(query string) ([]db.TechnicalService, error) {
	return t.technicalServiceStore.Search(query)

}
