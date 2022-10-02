package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type deviceTypeStore struct {
	db *gorm.DB
}

type DeviceTypeStore interface {
	Migration()
	Create(deviceType *expr_service.DeviceType) error
	Update(deviceType *expr_service.DeviceType) error
	Delete(deviceType *expr_service.DeviceType) error
	FindAll() ([]expr_service.DeviceType, error)
	FindAllByActive() ([]expr_service.DeviceType, error)
	FindByID(id int) (expr_service.DeviceType, error)
	FindBy(column string, value interface{}) ([]expr_service.DeviceType, error)
	Search(query string) ([]expr_service.DeviceType, error)
	Seed() error
}

func NewDeviceTypeStore(db *gorm.DB) DeviceTypeStore {
	return &deviceTypeStore{db: db}
}

func (d *deviceTypeStore) Migration() {
	d.db.AutoMigrate(&expr_service.DeviceType{})
}

func (d *deviceTypeStore) Create(deviceType *expr_service.DeviceType) error {
	return d.db.Create(deviceType).Error
}

func (d *deviceTypeStore) Update(deviceType *expr_service.DeviceType) error {
	return d.db.Save(deviceType).Error
}

func (d *deviceTypeStore) Delete(deviceType *expr_service.DeviceType) error {
	return d.db.Delete(deviceType).Error
}

func (d *deviceTypeStore) FindAll() ([]expr_service.DeviceType, error) {
	var deviceTypes []expr_service.DeviceType
	err := d.db.Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindAllByActive() ([]expr_service.DeviceType, error) {
	var deviceTypes []expr_service.DeviceType
	err := d.db.Where("deleted_at is null AND is_active = ?", true).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) FindByID(id int) (expr_service.DeviceType, error) {
	var deviceType expr_service.DeviceType
	err := d.db.First(&deviceType, id).Error
	return deviceType, err
}

func (d *deviceTypeStore) FindBy(column string, value interface{}) ([]expr_service.DeviceType, error) {
	var deviceTypes []expr_service.DeviceType
	err := d.db.Where(column+" = ?", value).Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Search(query string) ([]expr_service.DeviceType, error) {
	var deviceTypes []expr_service.DeviceType
	err := d.db.Where("name LIKE ?", "%"+query+"%").Find(&deviceTypes).Error
	return deviceTypes, err
}

func (d *deviceTypeStore) Seed() error {
	return nil
}
