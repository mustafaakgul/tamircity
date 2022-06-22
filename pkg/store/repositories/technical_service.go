package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type technicalServiceStore struct {
	db *gorm.DB
}

type TechnicalServiceStore interface {
	Create(model *db.TechnicalService) error
	Update(model *db.TechnicalService) error
	Delete(model *db.TechnicalService) error
	FindAll() ([]db.TechnicalService, error)
	FindByID(id int) (db.TechnicalService, error)
	FindBy(column string, value interface{}) ([]db.TechnicalService, error)
	Search(query string) ([]db.TechnicalService, error)
}

func NewTechnicalServiceStore(db *gorm.DB) TechnicalServiceStore {
	return &technicalServiceStore{db: db}
}

func (t technicalServiceStore) Create(model *db.TechnicalService) error {
	return t.db.Create(model).Error
}

func (t technicalServiceStore) Update(model *db.TechnicalService) error {
	return t.db.Save(model).Error
}

func (t technicalServiceStore) Delete(model *db.TechnicalService) error {
	return t.db.Delete(model).Error
}

func (t technicalServiceStore) FindAll() ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Find(&models).Error
	return models, err
}

func (t technicalServiceStore) FindByID(id int) (db.TechnicalService, error) {
	var model db.TechnicalService
	err := t.db.First(&model, id).Error
	return model, err
}

func (t technicalServiceStore) FindBy(column string, value interface{}) ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (t technicalServiceStore) Search(query string) ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
