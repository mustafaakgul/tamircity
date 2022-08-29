package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type brandStore struct {
	db *gorm.DB
}

//interface
type BrandStore interface {
	Migration()
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

func NewBrandStore(db *gorm.DB) *brandStore {
	return &brandStore{db: db}
}

func (b *brandStore) Migration() {
	b.db.AutoMigrate(&db.Brand{})
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
	var brands []db.Brand
	err := b.db.Model(&brands).Joins("inner join device_types_brands on brands.id = device_types_brands.brand_id").Where("device_types_brands.device_type_id = ?", deviceTypeId).Find(&brands).Error
	return brands, err
}

func (b *brandStore) Seed() error {
	brands := []*db.Brand{
		{Name: "Samsung", IsActive: true},
		{Name: "Apple", IsActive: true},
		{Name: "Xiaomi", IsActive: true},
		{Name: "Huawei", IsActive: true},
	}
	for _, brand := range brands {
		if err := b.db.Create(&brand).Error; err != nil {
			return err
		}
	}

	return nil
}
