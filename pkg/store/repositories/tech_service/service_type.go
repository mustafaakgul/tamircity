package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
)

type serviceTypeStore struct {
	db *gorm.DB
}

type ServiceTypeStore interface {
	Create(model *tech_service.ServiceType) error
	Update(model *tech_service.ServiceType) error
	Delete(model *tech_service.ServiceType) error
	FindAll() ([]tech_service.ServiceType, error)
	FindByID(id int) (tech_service.ServiceType, error)
	FindBy(column string, value interface{}) ([]tech_service.ServiceType, error)
	Search(query string) ([]tech_service.ServiceType, error)
	Seed() error
}

func NewServiceTypeStore(db *gorm.DB) ServiceTypeStore {
	return &serviceTypeStore{db: db}
}

func (s *serviceTypeStore) Create(model *tech_service.ServiceType) error {
	return s.db.Create(model).Error
}

func (s *serviceTypeStore) Update(model *tech_service.ServiceType) error {
	return s.db.Save(model).Error
}

func (s *serviceTypeStore) Delete(model *tech_service.ServiceType) error {
	return s.db.Delete(model).Error
}

func (s *serviceTypeStore) FindAll() ([]tech_service.ServiceType, error) {
	var models []tech_service.ServiceType
	err := s.db.Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) FindByID(id int) (tech_service.ServiceType, error) {
	var model tech_service.ServiceType
	err := s.db.First(&model, id).Error
	return model, err
}

func (s *serviceTypeStore) FindBy(column string, value interface{}) ([]tech_service.ServiceType, error) {
	var models []tech_service.ServiceType
	err := s.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Search(query string) ([]tech_service.ServiceType, error) {
	var models []tech_service.ServiceType
	err := s.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Seed() error {
	s.db.Create(&tech_service2.ServiceType1)
	s.db.Create(&tech_service2.ServiceType2)
	s.db.Create(&tech_service2.ServiceType3)
	s.db.Create(&tech_service2.ServiceType4)

	return nil
}
