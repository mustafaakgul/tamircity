package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type brandStore struct {
	db *gorm.DB
}

type BrandStore interface {
	Create(model *db.Brand) error
	Update(model *db.Brand) error
	Delete(model *db.Brand) error
	FindAll() ([]db.Brand, error)
	FindByID(id int) (db.Brand, error)
	FindBy(column string, value interface{}) ([]db.Brand, error)
	FindByDeviceTypeId(deviceTypeId int) ([]db.Brand, error)
	Search(query string) ([]db.Brand, error)
	Seed() error
}

func NewBrandStore(db *gorm.DB) BrandStore {
	return &brandStore{db: db}
}

func (b *brandStore) Create(model *db.Brand) error {
	return b.db.Create(model).Error
}

func (b *brandStore) Update(model *db.Brand) error {
	return b.db.Save(model).Error
}

func (b *brandStore) Delete(model *db.Brand) error {
	return b.db.Delete(model).Error
}

func (b *brandStore) FindAll() ([]db.Brand, error) {
	var models []db.Brand
	err := b.db.Find(&models).Error
	return models, err
}

func (b *brandStore) FindByID(id int) (db.Brand, error) {
	var model db.Brand
	err := b.db.First(&model, id).Error
	return model, err
}

func (b *brandStore) FindBy(column string, value interface{}) ([]db.Brand, error) {
	var models []db.Brand
	err := b.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (b *brandStore) Search(query string) ([]db.Brand, error) {
	var models []db.Brand
	err := b.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (b *brandStore) FindByDeviceTypeId(deviceTypeId int) ([]db.Brand, error) {
	//var brands []db.Brand
	//err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//return brands, err

	var brands []db.Brand
	err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//err := b.db.Model(brands).Preload("DeviceTypes").Where("device_types.id = ?", deviceTypeId).Find(&brands).Error
	return brands, err
}

func (b *brandStore) Seed() error {
	return nil
}
