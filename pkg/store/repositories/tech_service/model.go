package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
)

type modelStore struct {
	db *gorm.DB
}

//interface
type ModelStore interface {
	Migration()
	Create(model *tech_service.Model) error
	Update(model *tech_service.Model) error
	Delete(model *tech_service.Model) error
	FindAll() ([]tech_service.Model, error)
	FindByID(id int) (tech_service.Model, error)
	FindBy(column string, value interface{}) ([]tech_service.Model, error)
	Where(brandId int, deviceTypeId int) ([]tech_service.Model, error)
	Search(query string) ([]tech_service.Model, error)
	Seed() error
}

func NewModelStore(db *gorm.DB) ModelStore {
	return &modelStore{db: db}
}

func (m *modelStore) Migration() {
	m.db.AutoMigrate(&tech_service.Model{})
}

func (m *modelStore) Create(model *tech_service.Model) error {
	return m.db.Create(model).Error
}

func (m *modelStore) Update(model *tech_service.Model) error {
	return m.db.Save(model).Error
}

func (m *modelStore) Delete(model *tech_service.Model) error {
	return m.db.Delete(model).Error
}

func (m *modelStore) FindAll() ([]tech_service.Model, error) {
	var models []tech_service.Model
	err := m.db.Find(&models).Error
	return models, err
}

func (m *modelStore) FindByID(id int) (tech_service.Model, error) {
	var model tech_service.Model
	err := m.db.First(&model, id).Error
	return model, err
}

func (m *modelStore) FindBy(column string, value interface{}) ([]tech_service.Model, error) {
	var models []tech_service.Model
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (m *modelStore) Where(brandId int, deviceTypeId int) ([]tech_service.Model, error) {
	var models []tech_service.Model
	/*err := m.db.Model(&models).
	Joins("inner join models_brands on models.id = models_brands.model_id").
	Where("models_brands.brand_id = ?", brandId).
	Joins("inner join models_device_types on models.id = models_device_types.model_id").
	Where("models_device_types.device_type_id = ?", deviceTypeId).
	Find(&models).Error*/
	err := m.db.Where("brand_id = ? AND device_type_id = ?", brandId, deviceTypeId).Find(&models).Error
	return models, err
}

func (m *modelStore) Search(query string) ([]tech_service.Model, error) {
	var models []tech_service.Model
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (m *modelStore) Seed() error {
	models := []*tech_service.Model{
		seed_data.ModelSamsungPC,
		seed_data.ModelApplePC,
		seed_data.ModelLenovoPC,
		seed_data.ModelHpPC,
		seed_data.ModelAsusPC,
		seed_data.ModelAcerPC,
		seed_data.ModelToshibaPC,
		seed_data.ModelDellPC,
		seed_data.ModelCasperPC,
		seed_data.ModelHuaweiPC,
		seed_data.ModelMSIPC,
		seed_data.ModelMicrosoftPC,
		seed_data.ModelMonsterPC,
		seed_data.ModelAsusRogPC,
		seed_data.ModelGpad7,
		seed_data.ModelGpad8,
		seed_data.ModelLGOther,
		seed_data.ModelGMTAB,
		seed_data.ModelGMTAB4,
		seed_data.ModelGMTAB5,
		seed_data.ModelGMTAB10,
		seed_data.ModelGMTAB8,
		seed_data.ModelIphone11,
		seed_data.ModelIphone12,
		seed_data.ModelIphone12Pro,
		seed_data.ModelSamsungGalaxyS7,
		seed_data.ModelSamsungGalaxyS9,
		seed_data.ModelLenovoM7,
		seed_data.ModelLenovoM8,
		seed_data.ModeliPad6,
		seed_data.ModeliPad9,
	}

	for _, model := range models {
		if err := m.db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}
