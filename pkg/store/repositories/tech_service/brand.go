package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
)

type brandStore struct {
	db *gorm.DB
}

// interface
type BrandStore interface {
	Create(model *tech_service.Brand) error
	Update(model *tech_service.Brand) error
	Delete(model *tech_service.Brand) error
	FindAll() ([]tech_service.Brand, error)
	FindByID(id int) (tech_service.Brand, error)
	FindBy(column string, value interface{}) ([]tech_service.Brand, error)
	FindByDeviceTypeId(deviceTypeId int) ([]tech_service.Brand, error)
	Search(query string) ([]tech_service.Brand, error)
	Seed() error
}

func NewBrandStore(db *gorm.DB) BrandStore {
	return &brandStore{db: db}
}

func (b *brandStore) Create(model *tech_service.Brand) error {
	return b.db.Create(model).Error
}

func (b *brandStore) Update(model *tech_service.Brand) error {
	return b.db.Save(model).Error
}

func (b *brandStore) Delete(model *tech_service.Brand) error {
	return b.db.Delete(model).Error
}

func (b *brandStore) FindAll() ([]tech_service.Brand, error) {
	var models []tech_service.Brand
	err := b.db.Find(&models).Error
	return models, err
}

func (b *brandStore) FindByID(id int) (tech_service.Brand, error) {
	var model tech_service.Brand
	err := b.db.First(&model, id).Error
	return model, err
}

func (b *brandStore) FindBy(column string, value interface{}) ([]tech_service.Brand, error) {
	var models []tech_service.Brand
	err := b.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (b *brandStore) Search(query string) ([]tech_service.Brand, error) {
	var models []tech_service.Brand
	err := b.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (b *brandStore) FindByDeviceTypeId(deviceTypeId int) ([]tech_service.Brand, error) {
	//var brands []db.Brand
	//err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//return brands, err

	var brands []tech_service.Brand
	err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	//err := b.db.Model(brands).Preload("DeviceTypes").Where("device_types.id = ?", deviceTypeId).Find(&brands).Error
	return brands, err
}

func (b *brandStore) Seed() error {
	b.db.Create(&tech_service2.BrandApple)
	b.db.Create(&tech_service2.BrandSamsung)
	b.db.Create(&tech_service2.BrandLenovo)
	b.db.Create(&tech_service2.BrandNokia)
	b.db.Create(&tech_service2.BrandOppo)
	b.db.Create(&tech_service2.BrandXiaomi)
	b.db.Create(&tech_service2.BrandHuawei)
	b.db.Create(&tech_service2.BrandAsus)
	b.db.Create(&tech_service2.BrandLG)
	b.db.Create(&tech_service2.BrandHp)
	b.db.Create(&tech_service2.BrandAcer)
	b.db.Create(&tech_service2.BrandToshiba)
	b.db.Create(&tech_service2.BrandSony)
	b.db.Create(&tech_service2.BrandMicrosoft)
	b.db.Create(&tech_service2.BrandDell)
	b.db.Create(&tech_service2.BrandCasper)
	b.db.Create(&tech_service2.BrandVestel)
	b.db.Create(&tech_service2.BrandMsi)
	b.db.Create(&tech_service2.BrandMonster)
	b.db.Create(&tech_service2.BrandPhilips)
	b.db.Create(&tech_service2.BrandAsusRog)
	b.db.Create(&tech_service2.BrandGeneralMobile)
	b.db.Create(&tech_service2.BrandHometech)
	b.db.Create(&tech_service2.BrandAlcatel)
	b.db.Create(&tech_service2.BrandReeder)
	b.db.Create(&tech_service2.BrandAmazon)
	b.db.Create(&tech_service2.BrandBlackberry)
	b.db.Create(&tech_service2.BrandGoogle)
	b.db.Create(&tech_service2.BrandHiking)
	b.db.Create(&tech_service2.BrandHonor)
	b.db.Create(&tech_service2.BrandHTC)
	b.db.Create(&tech_service2.BrandInfinix)
	b.db.Create(&tech_service2.BrandOnePlus)
	b.db.Create(&tech_service2.BrandRealme)
	b.db.Create(&tech_service2.BrandTCL)
	b.db.Create(&tech_service2.BrandTecno)
	b.db.Create(&tech_service2.BrandVivo)

	return nil
}
