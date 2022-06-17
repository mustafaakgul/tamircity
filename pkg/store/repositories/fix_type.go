package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type fixTypeRepository struct {
	db *gorm.DB
}

//interface
type FixTypeRepository interface {
	Migration()
	Create(model *db.FixType) error
	Update(model *db.FixType) error
	Delete(model *db.FixType) error
	FindAll() ([]db.FixType, error)
	FindByID(id int) (db.FixType, error)
	FindBy(column string, value interface{}) ([]db.FixType, error)
	Search(query string) ([]db.FixType, error)
}

var _ FixTypeRepository = &fixTypeRepository{}

func NewFixTypeRepository(db *gorm.DB) *fixTypeRepository {
	return &fixTypeRepository{db: db}
}
func (m *fixTypeRepository) Migration() {
	m.db.AutoMigrate(&db.FixType{})
}
func (m *fixTypeRepository) Create(model *db.FixType) error {
	return m.db.Create(model).Error
}
func (m *fixTypeRepository) Update(model *db.FixType) error {
	return m.db.Save(model).Error
}
func (m *fixTypeRepository) Delete(model *db.FixType) error {
	return m.db.Delete(model).Error
}
func (m *fixTypeRepository) FindAll() ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Find(&models).Error
	return models, err
}
func (m *fixTypeRepository) FindByID(id int) (db.FixType, error) {
	var model db.FixType
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *fixTypeRepository) FindBy(column string, value interface{}) ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *fixTypeRepository) Search(query string) ([]db.FixType, error) {
	var models []db.FixType
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
