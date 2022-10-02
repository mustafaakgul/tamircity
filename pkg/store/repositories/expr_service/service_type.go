package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type serviceTypeStore struct {
	db *gorm.DB
}

type ServiceTypeStore interface {
	Create(model *expr_service.ServiceType) error
	Update(model *expr_service.ServiceType) error
	Delete(model *expr_service.ServiceType) error
	FindAll() ([]expr_service.ServiceType, error)
	FindByID(id int) (expr_service.ServiceType, error)
	FindBy(column string, value interface{}) ([]expr_service.ServiceType, error)
	Search(query string) ([]expr_service.ServiceType, error)
	Seed() error
}

func NewServiceTypeStore(db *gorm.DB) ServiceTypeStore {
	return &serviceTypeStore{db: db}
}

func (s *serviceTypeStore) Create(model *expr_service.ServiceType) error {
	return s.db.Create(model).Error
}

func (s *serviceTypeStore) Update(model *expr_service.ServiceType) error {
	return s.db.Save(model).Error
}

func (s *serviceTypeStore) Delete(model *expr_service.ServiceType) error {
	return s.db.Delete(model).Error
}

func (s *serviceTypeStore) FindAll() ([]expr_service.ServiceType, error) {
	var models []expr_service.ServiceType
	err := s.db.Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) FindByID(id int) (expr_service.ServiceType, error) {
	var model expr_service.ServiceType
	err := s.db.First(&model, id).Error
	return model, err
}

func (s *serviceTypeStore) FindBy(column string, value interface{}) ([]expr_service.ServiceType, error) {
	var models []expr_service.ServiceType
	err := s.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Search(query string) ([]expr_service.ServiceType, error) {
	var models []expr_service.ServiceType
	err := s.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Seed() error {
	return nil
}
