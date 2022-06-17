package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type extraServiceService struct {
	extraServiceStore repositories.ExtraServiceStore
}

type ExtraServiceService interface {
	Create(model *db.ExtraService) error
	Update(model *db.ExtraService) error
	Delete(model *db.ExtraService) error
	FindAll() ([]db.ExtraService, error)
	FindByID(id int) (db.ExtraService, error)
	FindBy(column string, value interface{}) ([]db.ExtraService, error)
	Search(query string) ([]db.ExtraService, error)
}

func NewExtraServiceService(extraServiceStore repositories.ExtraServiceStore) ExtraServiceService {
	return &extraServiceService{
		extraServiceStore: extraServiceStore,
	}
}

func (e *extraServiceService) Create(model *db.ExtraService) error {
	return e.extraServiceStore.Create(model)
}

func (e *extraServiceService) Update(model *db.ExtraService) error {
	return e.extraServiceStore.Update(model)
}

func (e *extraServiceService) Delete(model *db.ExtraService) error {
	return e.extraServiceStore.Delete(model)
}

func (e *extraServiceService) FindAll() ([]db.ExtraService, error) {
	return e.extraServiceStore.FindAll()
}

func (e *extraServiceService) FindByID(id int) (db.ExtraService, error) {
	return e.extraServiceStore.FindByID(id)
}

func (e *extraServiceService) FindBy(column string, value interface{}) ([]db.ExtraService, error) {
	return e.extraServiceStore.FindBy(column, value)
}

func (e *extraServiceService) Search(query string) ([]db.ExtraService, error) {
	return e.extraServiceStore.Search(query)
}
