package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
)

type extraServiceStore struct {
	db *gorm.DB
}

type ExtraServiceStore interface {
	Create(model *tech_service.ExtraService) error
	Update(model *tech_service.ExtraService) error
	Delete(model *tech_service.ExtraService) error
	FindAll() ([]tech_service.ExtraService, error)
	FindByID(id int) (tech_service.ExtraService, error)
	FindBy(column string, value interface{}) ([]tech_service.ExtraService, error)
	Search(query string) ([]tech_service.ExtraService, error)
	Seed() error
}

func NewExtraServiceStore(db *gorm.DB) ExtraServiceStore {
	return &extraServiceStore{db: db}
}

func (e *extraServiceStore) Create(model *tech_service.ExtraService) error {
	return e.db.Create(model).Error
}

func (e *extraServiceStore) Update(model *tech_service.ExtraService) error {
	return e.db.Save(model).Error
}

func (e *extraServiceStore) Delete(model *tech_service.ExtraService) error {
	return e.db.Delete(model).Error
}

func (e *extraServiceStore) FindAll() ([]tech_service.ExtraService, error) {
	var models []tech_service.ExtraService
	err := e.db.Find(&models).Error
	return models, err
}

func (e *extraServiceStore) FindByID(id int) (tech_service.ExtraService, error) {
	var model tech_service.ExtraService
	err := e.db.First(&model, id).Error
	return model, err
}

func (e *extraServiceStore) FindBy(column string, value interface{}) ([]tech_service.ExtraService, error) {
	var models []tech_service.ExtraService
	err := e.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (e *extraServiceStore) Search(query string) ([]tech_service.ExtraService, error) {
	var models []tech_service.ExtraService
	err := e.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
func (e *extraServiceStore) Seed() error {
	e.db.Create(&tech_service2.ExtraServiceFirst)
	e.db.Create(&tech_service2.ExtraServiceSecond)
	e.db.Create(&tech_service2.ExtraServiceThird)
	e.db.Create(&tech_service2.ExtraServiceFourth)

	return nil
}
