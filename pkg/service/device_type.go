package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web/tech_service"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type DeviceTypeService interface {
	Create(model *db.DeviceType) error
	Update(model *db.DeviceType) error
	Delete(model *db.DeviceType) error
	FindAll() ([]db.DeviceType, error)
	FindAllByActive() (response []tech_service.DeviceTypeResponse, err error)
	FindByID(id int) (db.DeviceType, error)
	FindBy(column string, value interface{}) ([]db.DeviceType, error)
	Search(query string) ([]db.DeviceType, error)
}

type deviceTypeService struct {
	deviceTypeStore repositories.DeviceTypeStore
}

func NewDeviceTypeService(deviceTypeStore repositories.DeviceTypeStore) DeviceTypeService {
	return &deviceTypeService{deviceTypeStore: deviceTypeStore}
}

func (m *deviceTypeService) Create(model *db.DeviceType) error {
	return m.deviceTypeStore.Create(model)
}

func (m *deviceTypeService) Update(model *db.DeviceType) error {
	return m.deviceTypeStore.Update(model)
}

func (m *deviceTypeService) Delete(model *db.DeviceType) error {
	return m.deviceTypeStore.Delete(model)
}

func (m *deviceTypeService) FindAll() ([]db.DeviceType, error) {
	return m.deviceTypeStore.FindAll()
}

func (m *deviceTypeService) FindAllByActive() (response []tech_service.DeviceTypeResponse, err error) {
	deviceTypes, err := m.deviceTypeStore.FindAllByActive()

	if err != nil {
		return nil, err
	}
	for _, deviceType := range deviceTypes {
		var deviceTypeResponse tech_service.DeviceTypeResponse
		deviceTypeResponse.Id = int(deviceType.ID)
		deviceTypeResponse.Name = deviceType.Name
		deviceTypeResponse.ShortDescription = deviceType.ShortDescription

		response = append(response, deviceTypeResponse)
	}

	return response, nil
}

func (m *deviceTypeService) FindByID(id int) (db.DeviceType, error) {
	return m.deviceTypeStore.FindByID(id)
}

func (m *deviceTypeService) FindBy(column string, value interface{}) ([]db.DeviceType, error) {
	return m.deviceTypeStore.FindBy(column, value)
}

func (m *deviceTypeService) Search(query string) ([]db.DeviceType, error) {
	return m.deviceTypeStore.Search(query)
}
