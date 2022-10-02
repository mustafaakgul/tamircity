package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
)

type fixTypeStore struct {
	db *gorm.DB
}

type FixTypeStore interface {
	Migration()
	Create(model *tech_service.FixType) error
	Update(model *tech_service.FixType) error
	Delete(model *tech_service.FixType) error
	FindAll() ([]tech_service.FixType, error)
	FindByID(id int) (tech_service.FixType, error)
	FindBy(column string, value interface{}) ([]tech_service.FixType, error)
	FindByDeviceTypeId(deviceTypeId int) ([]tech_service.FixType, error)
	Search(query string) ([]tech_service.FixType, error)
	Seed() error
}

func NewFixTypeStore(db *gorm.DB) FixTypeStore {
	return &fixTypeStore{db: db}
}

func (f *fixTypeStore) Migration() {
	f.db.AutoMigrate(&tech_service.FixType{})
}

func (f *fixTypeStore) Create(model *tech_service.FixType) error {
	return f.db.Create(model).Error
}

func (f *fixTypeStore) Update(model *tech_service.FixType) error {
	return f.db.Save(model).Error
}

func (f *fixTypeStore) Delete(model *tech_service.FixType) error {
	return f.db.Delete(model).Error
}

func (f *fixTypeStore) FindAll() ([]tech_service.FixType, error) {
	var models []tech_service.FixType
	err := f.db.Find(&models).Error
	return models, err
}

func (f *fixTypeStore) FindByID(id int) (tech_service.FixType, error) {
	var model tech_service.FixType
	err := f.db.First(&model, id).Error
	return model, err
}

func (f *fixTypeStore) FindBy(column string, value interface{}) ([]tech_service.FixType, error) {
	var models []tech_service.FixType
	err := f.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (f *fixTypeStore) FindByDeviceTypeId(deviceTypeId int) ([]tech_service.FixType, error) {
	var fixTypes []tech_service.FixType
	err := f.db.Model(&fixTypes).Joins("INNER JOIN device_types_fix_types on fix_types.id = device_types_fix_types.fix_type_id").Where("device_types_fix_types.device_type_id = ?", deviceTypeId).Find(&fixTypes).Error
	return fixTypes, err
}

func (f *fixTypeStore) Search(query string) ([]tech_service.FixType, error) {
	var models []tech_service.FixType
	err := f.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (f *fixTypeStore) Seed() error {

	fixTypes := []*tech_service.FixType{tech_service2.FixType1, tech_service2.FixType2, tech_service2.FixType3, tech_service2.FixType4}

	for _, fixType := range fixTypes {
		if err := f.db.Create(&fixType).Error; err != nil {
			return err
		}
	}

	return nil
}
