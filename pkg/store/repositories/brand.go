package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type brandRepository struct {
	db *gorm.DB
}

//interface
type BrandRepository interface {
	Migration()
	Create(model *db.Brand) error
	Update(model *db.Brand) error
	Delete(model *db.Brand) error
	FindAll() ([]db.Brand, error)
	FindByID(id int) (db.Brand, error)
	FindBy(column string, value interface{}) ([]db.Brand, error)
	Search(query string) ([]db.Brand, error)
}

var _ BrandRepository = &brandRepository{}

func NewBrandRepository(db *gorm.DB) *brandRepository {
	return &brandRepository{db: db}
}
func (m *brandRepository) Migration() {
	m.db.AutoMigrate(&db.Brand{})
}
func (m *brandRepository) Create(model *db.Brand) error {
	return m.db.Create(model).Error
}
func (m *brandRepository) Update(model *db.Brand) error {
	return m.db.Save(model).Error
}
func (m *brandRepository) Delete(model *db.Brand) error {
	return m.db.Delete(model).Error
}
func (m *brandRepository) FindAll() ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Find(&models).Error
	return models, err
}
func (m *brandRepository) FindByID(id int) (db.Brand, error) {
	var model db.Brand
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *brandRepository) FindBy(column string, value interface{}) ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *brandRepository) Search(query string) ([]db.Brand, error) {
	var models []db.Brand
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
