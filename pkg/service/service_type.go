package service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/models/web"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
)

type serviceTypeService struct {
	serviceTypeStore tech_service2.ServiceTypeStore
}

type ServiceTypeService interface {
	Create(model *tech_service.ServiceType) error
	Update(model *tech_service.ServiceType) error
	Delete(model *tech_service.ServiceType) error
	FindAll() ([]web.ServiceTypeResponse, error)
	FindByID(id int) (tech_service.ServiceType, error)
	FindBy(column string, value interface{}) ([]tech_service.ServiceType, error)
	Search(query string) ([]tech_service.ServiceType, error)
}

func NewServiceTypeService(serviceTypeStore tech_service2.ServiceTypeStore) ServiceTypeService {
	return &serviceTypeService{
		serviceTypeStore: serviceTypeStore,
	}
}

func (s *serviceTypeService) Create(model *tech_service.ServiceType) error {
	return s.serviceTypeStore.Create(model)
}

func (s *serviceTypeService) Update(model *tech_service.ServiceType) error {
	return s.serviceTypeStore.Update(model)
}

func (s *serviceTypeService) Delete(model *tech_service.ServiceType) error {
	return s.serviceTypeStore.Delete(model)
}

func (s *serviceTypeService) FindAll() (res []web.ServiceTypeResponse, err error) {
	serviceTypes, err := s.serviceTypeStore.FindAll()
	if err != nil {
		return nil, err
	}

	for _, serviceType := range serviceTypes {
		res = append(res, web.ServiceTypeResponse{Id: serviceType.ID, Name: serviceType.Name, Description: serviceType.Description, Price: serviceType.Price})
	}
	return res, nil
}

func (s *serviceTypeService) FindByID(id int) (tech_service.ServiceType, error) {
	return s.serviceTypeStore.FindByID(id)
}

func (s *serviceTypeService) FindBy(column string, value interface{}) ([]tech_service.ServiceType, error) {
	return s.serviceTypeStore.FindBy(column, value)
}

func (s *serviceTypeService) Search(query string) ([]tech_service.ServiceType, error) {
	return s.serviceTypeStore.Search(query)

}
