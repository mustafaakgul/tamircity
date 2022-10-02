package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type deviceTypeStore struct {
	db *gorm.DB
}

type DeviceTypeStore interface {
	Migration()
	Create(deviceType *db.DeviceType) error
	Update(deviceType *db.DeviceType) error
	Delete(deviceType *db.DeviceType) error
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

func (d *deviceTypeStore) Migration() {
	d.db.AutoMigrate(&db.DeviceType{})
}

func (d *deviceTypeStore) Create(deviceType *db.DeviceType) error {
	return d.db.Create(deviceType).Error
}

func (d *deviceTypeStore) Update(deviceType *db.DeviceType) error {
	return d.db.Save(deviceType).Error
}

func (d *deviceTypeStore) Delete(deviceType *db.DeviceType) error {
	return d.db.Delete(deviceType).Error
}

func (d *deviceTypeStore) FindAll() ([]db.DeviceType, error) {
	var deviceTypes []db.DeviceType
	err := d.db.Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindAllByActive() ([]db.DeviceType, error) {
	var deviceTypes []db.DeviceType
	err := d.db.Where("deleted_at is null AND is_active = ?", true).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindByID(id int) (db.DeviceType, error) {
	var deviceType db.DeviceType
	err := d.db.First(&deviceType, id).Error
	return deviceType, err
}

func (d *deviceTypeStore) FindBy(column string, value interface{}) ([]db.DeviceType, error) {
	var deviceTypes []db.DeviceType
	err := d.db.Where(column+" = ?", value).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Search(query string) ([]db.DeviceType, error) {
	var deviceTypes []db.DeviceType
	err := d.db.Where("name LIKE ?", "%"+query+"%").Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Seed() error {
	return nil
}
