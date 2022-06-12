package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type deviceTypeRepository struct {
	db *gorm.DB
}

//interface
type DeviceTypeRepository interface {
	Migration()
	Create(model *db.DeviceType) error
	Update(model *db.DeviceType) error
	Delete(model *db.DeviceType) error
	FindAll() ([]db.DeviceType, error)
	FindByID(id int) (db.DeviceType, error)
	FindBy(column string, value interface{}) ([]db.DeviceType, error)
	Search(query string) ([]db.DeviceType, error)
}

var _ DeviceTypeRepository = &deviceTypeRepository{}

func NewDeviceTypeRepository(db *gorm.DB) *deviceTypeRepository {
	return &deviceTypeRepository{db: db}
}
func (m *deviceTypeRepository) Migration() {
	m.db.AutoMigrate(&db.DeviceType{})
}
func (m *deviceTypeRepository) Create(model *db.DeviceType) error {
	return m.db.Create(model).Error
}
func (m *deviceTypeRepository) Update(model *db.DeviceType) error {
	return m.db.Save(model).Error
}
func (m *deviceTypeRepository) Delete(model *db.DeviceType) error {
	return m.db.Delete(model).Error
}
func (m *deviceTypeRepository) FindAll() ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Find(&models).Error
	return models, err
}
func (m *deviceTypeRepository) FindByID(id int) (db.DeviceType, error) {
	var model db.DeviceType
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *deviceTypeRepository) FindBy(column string, value interface{}) ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *deviceTypeRepository) Search(query string) ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
