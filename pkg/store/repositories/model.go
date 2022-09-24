package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
)

type modelStore struct {
	db *gorm.DB
}

//interface
type ModelStore interface {
	Migration()
	Create(model *db.Model) error
	Update(model *db.Model) error
	Delete(model *db.Model) error
	FindAll() ([]db.Model, error)
	FindByID(id int) (db.Model, error)
	FindBy(column string, value interface{}) ([]db.Model, error)
	Where(brandId int, deviceTypeId int) ([]db.Model, error)
	Search(query string) ([]db.Model, error)
	Seed() error
}

func NewModelStore(db *gorm.DB) ModelStore {
	return &modelStore{db: db}
}

func (m *modelStore) Migration() {
	m.db.AutoMigrate(&db.Model{})
}

func (m *modelStore) Create(model *db.Model) error {
	return m.db.Create(model).Error
}

func (m *modelStore) Update(model *db.Model) error {
	return m.db.Save(model).Error
}

func (m *modelStore) Delete(model *db.Model) error {
	return m.db.Delete(model).Error
}

func (m *modelStore) FindAll() ([]db.Model, error) {
	var models []db.Model
	err := m.db.Find(&models).Error
	return models, err
}

func (m *modelStore) FindByID(id int) (db.Model, error) {
	var model db.Model
	err := m.db.First(&model, id).Error
	return model, err
}

func (m *modelStore) FindBy(column string, value interface{}) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (m *modelStore) Where(brandId int, deviceTypeId int) ([]db.Model, error) {
	var models []db.Model
	/*err := m.db.Model(&models).
	Joins("inner join models_brands on models.id = models_brands.model_id").
	Where("models_brands.brand_id = ?", brandId).
	Joins("inner join models_device_types on models.id = models_device_types.model_id").
	Where("models_device_types.device_type_id = ?", deviceTypeId).
	Find(&models).Error*/
	err := m.db.Where("brand_id = ? AND device_type_id = ?", brandId, deviceTypeId).Find(&models).Error
	return models, err
}

func (m *modelStore) Search(query string) ([]db.Model, error) {
	var models []db.Model
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (m *modelStore) Seed() error {
	models := []*db.Model{seed_data.ModelIphone11, seed_data.ModelIphone12, seed_data.ModelIphone12Pro, seed_data.ModelSamsungGalaxyS7, seed_data.ModelSamsungGalaxyS9, seed_data.ModelLenovoM7, seed_data.ModelLenovoM8, seed_data.ModeliPad6, seed_data.ModeliPad9}

	for _, model := range models {
		if err := m.db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}
