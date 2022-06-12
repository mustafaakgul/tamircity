package service

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
)

type DeviceTypeService interface {
	Create(model *db.DeviceType) error
	Update(model *db.DeviceType) error
	Delete(model *db.DeviceType) error
	FindAll() ([]db.DeviceType, error)
	FindByID(id int) (db.DeviceType, error)
	FindBy(column string, value interface{}) ([]db.DeviceType, error)
	Search(query string) ([]db.DeviceType, error)
}

type deviceTypeService struct {
	deviveTypeRepository repositories.DeviceTypeRepository
}

func NewDeviceTypeService(deviveTypeRepository repositories.DeviceTypeRepository) DeviceTypeService {
	return &deviceTypeService{deviveTypeRepository: deviveTypeRepository}
}

func (m *deviceTypeService) Create(model *db.DeviceType) error {
	return m.deviveTypeRepository.Create(model)
}
func (m *deviceTypeService) Update(model *db.DeviceType) error {
	return m.deviveTypeRepository.Update(model)
}
func (m *deviceTypeService) Delete(model *db.DeviceType) error {
	return m.deviveTypeRepository.Delete(model)
}
func (m *deviceTypeService) FindAll() ([]db.DeviceType, error) {
	return m.deviveTypeRepository.FindAll()
}
func (m *deviceTypeService) FindByID(id int) (db.DeviceType, error) {
	return m.deviveTypeRepository.FindByID(id)
}
func (m *deviceTypeService) FindBy(column string, value interface{}) ([]db.DeviceType, error) {
	return m.deviveTypeRepository.FindBy(column, value)
}
func (m *deviceTypeService) Search(query string) ([]db.DeviceType, error) {
	return m.deviveTypeRepository.Search(query)
}
