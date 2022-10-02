package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type DeviceTypeService interface {
	Create(model *tech_service.DeviceType) error
	Update(model *tech_service.DeviceType) error
	Delete(model *tech_service.DeviceType) error
	FindAll() ([]tech_service.DeviceType, error)
	FindAllByActive() (response []web.DeviceTypeResponse, err error)
	FindByID(id int) (tech_service.DeviceType, error)
	FindBy(column string, value interface{}) ([]tech_service.DeviceType, error)
	Search(query string) ([]tech_service.DeviceType, error)
}

type deviceTypeService struct {
	deviceTypeStore tech_service2.DeviceTypeStore
}

func NewDeviceTypeService(deviceTypeStore tech_service2.DeviceTypeStore) DeviceTypeService {
	return &deviceTypeService{deviceTypeStore: deviceTypeStore}
}

func (m *deviceTypeService) Create(model *tech_service.DeviceType) error {
	return m.deviceTypeStore.Create(model)
}

func (m *deviceTypeService) Update(model *tech_service.DeviceType) error {
	return m.deviceTypeStore.Update(model)
}

func (m *deviceTypeService) Delete(model *tech_service.DeviceType) error {
	return m.deviceTypeStore.Delete(model)
}

func (m *deviceTypeService) FindAll() ([]tech_service.DeviceType, error) {
	return m.deviceTypeStore.FindAll()
}

func (m *deviceTypeService) FindAllByActive() (response []web.DeviceTypeResponse, err error) {
	deviceTypes, err := m.deviceTypeStore.FindAllByActive()

	if err != nil {
		return nil, err
	}
	for _, deviceType := range deviceTypes {
		var deviceTypeResponse web.DeviceTypeResponse
		deviceTypeResponse.Id = int(deviceType.ID)
		deviceTypeResponse.Name = deviceType.Name
		deviceTypeResponse.ShortDescription = deviceType.ShortDescription

		response = append(response, deviceTypeResponse)
	}

	return response, nil
}

func (m *deviceTypeService) FindByID(id int) (tech_service.DeviceType, error) {
	return m.deviceTypeStore.FindByID(id)
}

func (m *deviceTypeService) FindBy(column string, value interface{}) ([]tech_service.DeviceType, error) {
	return m.deviceTypeStore.FindBy(column, value)
}

func (m *deviceTypeService) Search(query string) ([]tech_service.DeviceType, error) {
	return m.deviceTypeStore.Search(query)
}
