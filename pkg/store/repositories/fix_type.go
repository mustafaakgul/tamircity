package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type fixTypeStore struct {
	db *gorm.DB
}

type FixTypeStore interface {
	Migration()
	Create(model *db.FixType) error
	Update(model *db.FixType) error
	Delete(model *db.FixType) error
	FindAll() ([]db.FixType, error)
	FindByID(id int) (db.FixType, error)
	FindBy(column string, value interface{}) ([]db.FixType, error)
	FindByDeviceTypeId(deviceTypeId int) ([]db.FixType, error)
	Search(query string) ([]db.FixType, error)
	Seed() error
}

func NewFixTypeStore(db *gorm.DB) FixTypeStore {
	return &fixTypeStore{db: db}
}

func (f *fixTypeStore) Migration() {
	f.db.AutoMigrate(&db.FixType{})
}

func (f *fixTypeStore) Create(model *db.FixType) error {
	return f.db.Create(model).Error
}

func (f *fixTypeStore) Update(model *db.FixType) error {
	return f.db.Save(model).Error
}

func (f *fixTypeStore) Delete(model *db.FixType) error {
	return f.db.Delete(model).Error
}

func (f *fixTypeStore) FindAll() ([]db.FixType, error) {
	var models []db.FixType
	err := f.db.Find(&models).Error
	return models, err
}

func (f *fixTypeStore) FindByID(id int) (db.FixType, error) {
	var model db.FixType
	err := f.db.First(&model, id).Error
	return model, err
}

func (f *fixTypeStore) FindBy(column string, value interface{}) ([]db.FixType, error) {
	var models []db.FixType
	err := f.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (f *fixTypeStore) FindByDeviceTypeId(deviceTypeId int) ([]db.FixType, error) {
	var fixTypes []db.FixType
	err := f.db.Model(&fixTypes).Joins("INNER JOIN device_types_fix_types on fix_types.id = device_types_fix_types.fix_type_id").Where("device_types_fix_types.device_type_id = ?", deviceTypeId).Find(&fixTypes).Error
	return fixTypes, err
}

func (f *fixTypeStore) Search(query string) ([]db.FixType, error) {
	var models []db.FixType
	err := f.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (f *fixTypeStore) Seed() error {

	fixTypes := []*db.FixType{fixType1, fixType2, fixType3, fixType4}

	for _, fixType := range fixTypes {
		if err := f.db.Create(&fixType).Error; err != nil {
			return err
		}
	}

	return nil
}
