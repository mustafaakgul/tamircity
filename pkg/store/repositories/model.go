package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type modelStore struct {
	db *gorm.DB
}

//interface
type ModelStore interface {
	Migration()
	Create(model *db.Model) error
	Update(model *db.Model) error
	Delete(model *db.Model) error
	FindAll() ([]db.Model, error)
	FindByID(id int) (db.Model, error)
	FindBy(column string, value interface{}) ([]db.Model, error)
	Where(query string) ([]db.Model, error)
	Search(query string) ([]db.Model, error)
}

func NewModelStore(db *gorm.DB) ModelStore {
	return &modelStore{db: db}
}

func (m *modelStore) Migration() {
	m.db.AutoMigrate(&db.Model{})
}

func (m *modelStore) Create(model *db.Model) error {
	return m.db.Create(model).Error
}

func (m *modelStore) Update(model *db.Model) error {
	return m.db.Save(model).Error
}

func (m *modelStore) Delete(model *db.Model) error {
	return m.db.Delete(model).Error
}

func (m *modelStore) FindAll() ([]db.Model, error) {
	var models []db.Model
	err := m.db.Find(&models).Error
	return models, err
}

func (m *modelStore) FindByID(id int) (db.Model, error) {
	var model db.Model
	err := m.db.First(&model, id).Error
	return model, err
}

func (m *modelStore) FindBy(column string, value interface{}) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (m *modelStore) Where(query string) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where(query).Find(&models).Error
	return models, err
}

func (m *modelStore) Search(query string) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
