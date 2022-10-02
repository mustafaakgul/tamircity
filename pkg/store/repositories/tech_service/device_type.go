package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
)

type deviceTypeStore struct {
	db *gorm.DB
}

type DeviceTypeStore interface {
	Migration()
	Create(deviceType *tech_service.DeviceType) error
	Update(deviceType *tech_service.DeviceType) error
	Delete(deviceType *tech_service.DeviceType) error
	FindAll() ([]tech_service.DeviceType, error)
	FindAllByActive() ([]tech_service.DeviceType, error)
	FindByID(id int) (tech_service.DeviceType, error)
	FindBy(column string, value interface{}) ([]tech_service.DeviceType, error)
	Search(query string) ([]tech_service.DeviceType, error)
	Seed() error
}

func NewDeviceTypeStore(db *gorm.DB) DeviceTypeStore {
	return &deviceTypeStore{db: db}
}

func (d *deviceTypeStore) Migration() {
	d.db.AutoMigrate(&tech_service.DeviceType{})
}

func (d *deviceTypeStore) Create(deviceType *tech_service.DeviceType) error {
	return d.db.Create(deviceType).Error
}

func (d *deviceTypeStore) Update(deviceType *tech_service.DeviceType) error {
	return d.db.Save(deviceType).Error
}

func (d *deviceTypeStore) Delete(deviceType *tech_service.DeviceType) error {
	return d.db.Delete(deviceType).Error
}

func (d *deviceTypeStore) FindAll() ([]tech_service.DeviceType, error) {
	var deviceTypes []tech_service.DeviceType
	err := d.db.Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindAllByActive() ([]tech_service.DeviceType, error) {
	var deviceTypes []tech_service.DeviceType
	err := d.db.Where("deleted_at is null AND is_active = ?", true).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindByID(id int) (tech_service.DeviceType, error) {
	var deviceType tech_service.DeviceType
	err := d.db.First(&deviceType, id).Error
	return deviceType, err
}

func (d *deviceTypeStore) FindBy(column string, value interface{}) ([]tech_service.DeviceType, error) {
	var deviceTypes []tech_service.DeviceType
	err := d.db.Where(column+" = ?", value).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Search(query string) ([]tech_service.DeviceType, error) {
	var deviceTypes []tech_service.DeviceType
	err := d.db.Where("name LIKE ?", "%"+query+"%").Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Seed() error {
	deviceTypes := []*tech_service.DeviceType{seed_data.DeviceTypePc, seed_data.DeviceTypePhone, seed_data.DeviceTypeTablet}
	for _, deviceType := range deviceTypes {
		if err := d.db.Create(&deviceType).Error; err != nil {
			return err
		}
	}
	return nil
}
