package service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web/tech_service"
	"github.com/anthophora/tamircity/pkg/store/repositories"
)

type serviceTypeService struct {
	serviceTypeStore repositories.ServiceTypeStore
}

type ServiceTypeService interface {
	Create(model *db.ServiceType) error
	Update(model *db.ServiceType) error
	Delete(model *db.ServiceType) error
	FindAll() ([]tech_service.ServiceTypeResponse, error)
	FindByID(id int) (db.ServiceType, error)
	FindBy(column string, value interface{}) ([]db.ServiceType, error)
	Search(query string) ([]db.ServiceType, error)
}

func NewServiceTypeService(serviceTypeStore repositories.ServiceTypeStore) ServiceTypeService {
	return &serviceTypeService{
		serviceTypeStore: serviceTypeStore,
	}
}

func (s *serviceTypeService) Create(model *db.ServiceType) error {
	return s.serviceTypeStore.Create(model)
}

func (s *serviceTypeService) Update(model *db.ServiceType) error {
	return s.serviceTypeStore.Update(model)
}

func (s *serviceTypeService) Delete(model *db.ServiceType) error {
	return s.serviceTypeStore.Delete(model)
}

func (s *serviceTypeService) FindAll() (res []tech_service.ServiceTypeResponse, err error) {
	serviceTypes, err := s.serviceTypeStore.FindAll()
	if err != nil {
		return nil, err
	}

	for _, serviceType := range serviceTypes {
		res = append(res, tech_service.ServiceTypeResponse{Id: serviceType.ID, Name: serviceType.Name, Description: serviceType.Description, Price: serviceType.Price})
	}
	return res, nil
}

func (s *serviceTypeService) FindByID(id int) (db.ServiceType, error) {
	return s.serviceTypeStore.FindByID(id)
}

func (s *serviceTypeService) FindBy(column string, value interface{}) ([]db.ServiceType, error) {
	return s.serviceTypeStore.FindBy(column, value)
}

func (s *serviceTypeService) Search(query string) ([]db.ServiceType, error) {
	return s.serviceTypeStore.Search(query)

}
