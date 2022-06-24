package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type brandStore struct {
	db *gorm.DB
}

//interface
type BrandStore interface {
	Migration()
	Create(model *db.Brand) error
	Update(model *db.Brand) error
	Delete(model *db.Brand) error
	FindAll() ([]db.Brand, error)
	FindByID(id int) (db.Brand, error)
	FindBy(column string, value interface{}) ([]db.Brand, error)
	Search(query string) ([]db.Brand, error)
}

func NewBrandStore(db *gorm.DB) *brandStore {
	return &brandStore{db: db}
}
func (m *brandStore) Migration() {
	m.db.AutoMigrate(&db.Brand{})
}
func (m *brandStore) Create(model *db.Brand) error {
	return m.db.Create(model).Error
}
func (m *brandStore) Update(model *db.Brand) error {
	return m.db.Save(model).Error
}
func (m *brandStore) Delete(model *db.Brand) error {
	return m.db.Delete(model).Error
}
func (m *brandStore) FindAll() ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Find(&models).Error
	return models, err
}
func (m *brandStore) FindByID(id int) (db.Brand, error) {
	var model db.Brand
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *brandStore) FindBy(column string, value interface{}) ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *brandStore) Search(query string) ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
