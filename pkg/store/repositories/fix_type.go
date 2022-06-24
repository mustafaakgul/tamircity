package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type fixTypeStore struct {
	db *gorm.DB
}

//interface
type FixTypeStore interface {
	Migration()
	Create(model *db.FixType) error
	Update(model *db.FixType) error
	Delete(model *db.FixType) error
	FindAll() ([]db.FixType, error)
	FindByID(id int) (db.FixType, error)
	FindBy(column string, value interface{}) ([]db.FixType, error)
	Search(query string) ([]db.FixType, error)
}

func NewFixTypeStore(db *gorm.DB) FixTypeStore {
	return &fixTypeStore{db: db}
}
func (m *fixTypeStore) Migration() {
	m.db.AutoMigrate(&db.FixType{})
}
func (m *fixTypeStore) Create(model *db.FixType) error {
	return m.db.Create(model).Error
}
func (m *fixTypeStore) Update(model *db.FixType) error {
	return m.db.Save(model).Error
}
func (m *fixTypeStore) Delete(model *db.FixType) error {
	return m.db.Delete(model).Error
}
func (m *fixTypeStore) FindAll() ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Find(&models).Error
	return models, err
}
func (m *fixTypeStore) FindByID(id int) (db.FixType, error) {
	var model db.FixType
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *fixTypeStore) FindBy(column string, value interface{}) ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *fixTypeStore) Search(query string) ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
