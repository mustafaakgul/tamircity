package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type brandStore struct {
	db *gorm.DB
}

type BrandStore interface {
	Create(model *expr_service.Brand) error
	Update(model *expr_service.Brand) error
	Delete(model *expr_service.Brand) error
	FindAll() ([]expr_service.Brand, error)
	FindByID(id int) (expr_service.Brand, error)
	FindBy(column string, value interface{}) ([]expr_service.Brand, error)
	FindByDeviceTypeId(deviceTypeId int) ([]expr_service.Brand, error)
	Search(query string) ([]expr_service.Brand, error)
	Seed() error
}

func NewBrandStore(db *gorm.DB) BrandStore {
	return &brandStore{db: db}
}

func (b *brandStore) Create(model *expr_service.Brand) error {
	return b.db.Create(model).Error
}

func (b *brandStore) Update(model *expr_service.Brand) error {
	return b.db.Save(model).Error
}

func (b *brandStore) Delete(model *expr_service.Brand) error {
	return b.db.Delete(model).Error
}

func (b *brandStore) FindAll() ([]expr_service.Brand, error) {
	var models []expr_service.Brand
	err := b.db.Find(&models).Error
	return models, err
}

func (b *brandStore) FindByID(id int) (expr_service.Brand, error) {
	var model expr_service.Brand
	err := b.db.First(&model, id).Error
	return model, err
}

func (b *brandStore) FindBy(column string, value interface{}) ([]expr_service.Brand, error) {
	var models []expr_service.Brand
	err := b.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (b *brandStore) Search(query string) ([]expr_service.Brand, error) {
	var models []expr_service.Brand
	err := b.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (b *brandStore) FindByDeviceTypeId(deviceTypeId int) ([]expr_service.Brand, error) {
	//var brands []db.Brand
	//err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//return brands, err

	var brands []expr_service.Brand
	err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//err := b.db.Model(brands).Preload("DeviceTypes").Where("device_types.id = ?", deviceTypeId).Find(&brands).Error
	return brands, err
}

func (b *brandStore) Seed() error {
	return nil
}
