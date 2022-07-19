package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type deviceTypeStore struct {
	db *gorm.DB
}

//interface
type DeviceTypeStore interface {
	Migration()
	Create(model *db.DeviceType) error
	Update(model *db.DeviceType) error
	Delete(model *db.DeviceType) error
	FindAll() ([]db.DeviceType, error)
	FindAllByActive() ([]db.DeviceType, error)
	FindByID(id int) (db.DeviceType, error)
	FindBy(column string, value interface{}) ([]db.DeviceType, error)
	Search(query string) ([]db.DeviceType, error)
	Seed() error
}

func NewDeviceTypeStore(db *gorm.DB) DeviceTypeStore {
	return &deviceTypeStore{db: db}
}
func (m *deviceTypeStore) Migration() {
	m.db.AutoMigrate(&db.DeviceType{})
}
func (m *deviceTypeStore) Create(model *db.DeviceType) error {
	return m.db.Create(model).Error
}
func (m *deviceTypeStore) Update(model *db.DeviceType) error {
	return m.db.Save(model).Error
}
func (m *deviceTypeStore) Delete(model *db.DeviceType) error {
	return m.db.Delete(model).Error
}
func (m *deviceTypeStore) FindAll() ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Find(&models).Error
	return models, err
}
func (m *deviceTypeStore) FindAllByActive() ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Where("deleted_at is null AND is_active = 1").Find(&models).Error
	return models, err
}
func (m *deviceTypeStore) FindByID(id int) (db.DeviceType, error) {
	var model db.DeviceType
	err := m.db.First(&model, id).Error
	return model, err
}
func (m *deviceTypeStore) FindBy(column string, value interface{}) ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}
func (m *deviceTypeStore) Search(query string) ([]db.DeviceType, error) {
	var models []db.DeviceType
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}
func (m *deviceTypeStore) Seed() error {
	deviceTypes = []*db.DeviceType{
		{Name: "Personel Computer", ShortDescription: "PC", IsActive: true},
		{Name: "Phone", ShortDescription: "PHONE", IsActive: true},
		{Name: "Tablet", ShortDescription: "TABLET", IsActive: true}
	}
	for _, deviceType := range deviceTypes {
		if err := b.db.Create(&deviceType).Error; err != nil {
			return err
		}
	}
	return nil
}
