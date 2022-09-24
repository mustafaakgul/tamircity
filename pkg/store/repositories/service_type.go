package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
)

type serviceTypeStore struct {
	db *gorm.DB
}

type ServiceTypeStore interface {
	Create(model *db.ServiceType) error
	Update(model *db.ServiceType) error
	Delete(model *db.ServiceType) error
	FindAll() ([]db.ServiceType, error)
	FindByID(id int) (db.ServiceType, error)
	FindBy(column string, value interface{}) ([]db.ServiceType, error)
	Search(query string) ([]db.ServiceType, error)
	Seed() error
}

func NewServiceTypeStore(db *gorm.DB) ServiceTypeStore {
	return &serviceTypeStore{db: db}
}

func (s *serviceTypeStore) Create(model *db.ServiceType) error {
	return s.db.Create(model).Error
}

func (s *serviceTypeStore) Update(model *db.ServiceType) error {
	return s.db.Save(model).Error
}

func (s *serviceTypeStore) Delete(model *db.ServiceType) error {
	return s.db.Delete(model).Error
}

func (s *serviceTypeStore) FindAll() ([]db.ServiceType, error) {
	var models []db.ServiceType
	err := s.db.Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) FindByID(id int) (db.ServiceType, error) {
	var model db.ServiceType
	err := s.db.First(&model, id).Error
	return model, err
}

func (s *serviceTypeStore) FindBy(column string, value interface{}) ([]db.ServiceType, error) {
	var models []db.ServiceType
	err := s.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Search(query string) ([]db.ServiceType, error) {
	var models []db.ServiceType
	err := s.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (s *serviceTypeStore) Seed() error {

	s.db.Create(&seed_data.ServiceType1)
	s.db.Create(&seed_data.ServiceType2)
	s.db.Create(&seed_data.ServiceType3)
	s.db.Create(&seed_data.ServiceType4)

	return nil
}
