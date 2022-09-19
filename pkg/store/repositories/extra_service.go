package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type extraServiceStore struct {
	db *gorm.DB
}

type ExtraServiceStore interface {
	Create(model *db.ExtraService) error
	Update(model *db.ExtraService) error
	Delete(model *db.ExtraService) error
	FindAll() ([]db.ExtraService, error)
	FindByID(id int) (db.ExtraService, error)
	FindBy(column string, value interface{}) ([]db.ExtraService, error)
	Search(query string) ([]db.ExtraService, error)
	Seed() error
}

func NewExtraServiceStore(db *gorm.DB) ExtraServiceStore {
	return &extraServiceStore{db: db}
}

func (e *extraServiceStore) Create(model *db.ExtraService) error {
	return e.db.Create(model).Error
}

func (e *extraServiceStore) Update(model *db.ExtraService) error {
	return e.db.Save(model).Error
}

func (e *extraServiceStore) Delete(model *db.ExtraService) error {
	return e.db.Delete(model).Error
}

func (e *extraServiceStore) FindAll() ([]db.ExtraService, error) {
	var models []db.ExtraService
	err := e.db.Find(&models).Error
	return models, err
}

func (e *extraServiceStore) FindByID(id int) (db.ExtraService, error) {
	var model db.ExtraService
	err := e.db.First(&model, id).Error
	return model, err
}

func (e *extraServiceStore) FindBy(column string, value interface{}) ([]db.ExtraService, error) {
	var models []db.ExtraService
	err := e.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (e *extraServiceStore) Search(query string) ([]db.ExtraService, error) {
	var models []db.ExtraService
	err := e.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
func (e *extraServiceStore) Seed() error {
	e.db.Create(&extraServiceFirst)
	e.db.Create(&extraServiceSecond)
	e.db.Create(&extraServiceThird)
	e.db.Create(&extraServiceFourth)

	return nil
}
