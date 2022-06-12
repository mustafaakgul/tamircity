package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type ModelService interface {
	Create(model *db.Model) error
	Update(model *db.Model) error
	Delete(model *db.Model) error
	FindAll() ([]db.Model, error)
	FindByID(id int) (db.Model, error)
	FindBy(column string, value interface{}) ([]db.Model, error)
	Search(query string) ([]db.Model, error)
}

type modelService struct {
	modelRepository repositories.ModelRepository
}

func NewModelService(modelRepository repositories.ModelRepository) ModelService {
	return &modelService{modelRepository: modelRepository}
}
func (m *modelService) Create(model *db.Model) error {
	return m.modelRepository.Create(model)
}
func (m *modelService) Update(model *db.Model) error {
	return m.modelRepository.Update(model)
}
func (m *modelService) Delete(model *db.Model) error {
	return m.modelRepository.Delete(model)
}
func (m *modelService) FindAll() ([]db.Model, error) {
	return m.modelRepository.FindAll()
}
func (m *modelService) FindByID(id int) (db.Model, error) {
	return m.modelRepository.FindByID(id)
}
func (m *modelService) FindBy(column string, value interface{}) ([]db.Model, error) {
	return m.modelRepository.FindBy(column, value)
}
func (m *modelService) Search(query string) ([]db.Model, error) {
	return m.modelRepository.Search(query)
}
