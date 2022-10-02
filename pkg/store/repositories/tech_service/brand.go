package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
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
	b.db.Create(&seed_data.BrandApple)
	b.db.Create(&seed_data.BrandSamsung)
	b.db.Create(&seed_data.BrandLenovo)
	b.db.Create(&seed_data.BrandNokia)
	b.db.Create(&seed_data.BrandOppo)
	b.db.Create(&seed_data.BrandXiaomi)
	b.db.Create(&seed_data.BrandHuawei)
	b.db.Create(&seed_data.BrandAsus)
	b.db.Create(&seed_data.BrandLG)
	b.db.Create(&seed_data.BrandHp)
	b.db.Create(&seed_data.BrandAcer)
	b.db.Create(&seed_data.BrandToshiba)
	b.db.Create(&seed_data.BrandSony)
	b.db.Create(&seed_data.BrandMicrosoft)
	b.db.Create(&seed_data.BrandDell)
	b.db.Create(&seed_data.BrandCasper)
	b.db.Create(&seed_data.BrandVestel)
	b.db.Create(&seed_data.BrandMsi)
	b.db.Create(&seed_data.BrandMonster)
	b.db.Create(&seed_data.BrandPhilips)
	b.db.Create(&seed_data.BrandAsusRog)
	b.db.Create(&seed_data.BrandGeneralMobile)
	b.db.Create(&seed_data.BrandHometech)
	b.db.Create(&seed_data.BrandAlcatel)
	b.db.Create(&seed_data.BrandReeder)
	b.db.Create(&seed_data.BrandAmazon)
	b.db.Create(&seed_data.BrandBlackberry)
	b.db.Create(&seed_data.BrandGoogle)
	b.db.Create(&seed_data.BrandHiking)
	b.db.Create(&seed_data.BrandHonor)
	b.db.Create(&seed_data.BrandHTC)
	b.db.Create(&seed_data.BrandInfinix)
	b.db.Create(&seed_data.BrandOnePlus)
	b.db.Create(&seed_data.BrandRealme)
	b.db.Create(&seed_data.BrandTCL)
	b.db.Create(&seed_data.BrandTecno)
	b.db.Create(&seed_data.BrandVivo)

	return nil
}
