package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type modelRepository struct {
	db *gorm.DB
}

//interface
type ModelRepository interface {
	Migration()
	Create(model *db.Model) error
	Update(model *db.Model) error
	Delete(model *db.Model) error
	FindAll() ([]db.Model, error)
	FindByID(id int) (db.Model, error)
	FindBy(column string, value interface{}) ([]db.Model, error)
	Search(query string) ([]db.Model, error)
}

var _ ModelRepository = &modelRepository{}

func NewModelRepository(db *gorm.DB) *modelRepository {
	return &modelRepository{db: db}
}
func (m *modelRepository) Migration() {
	m.db.AutoMigrate(&db.Model{})
}
func (m *modelRepository) Create(model *db.Model) error {
	return m.db.Create(model).Error
}
func (m *modelRepository) Update(model *db.Model) error {
	return m.db.Save(model).Error
}
func (m *modelRepository) Delete(model *db.Model) error {
	return m.db.Delete(model).Error
}
func (m *modelRepository) FindAll() ([]db.Model, error) {
	var models []db.Model
	err := m.db.Find(&models).Error
	return models, err
}
func (m *modelRepository) FindByID(id int) (db.Model, error) {
	var model db.Model
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *modelRepository) FindBy(column string, value interface{}) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *modelRepository) Search(query string) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
