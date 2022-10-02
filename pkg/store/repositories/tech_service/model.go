package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
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
		tech_service2.ModelSamsungPC,
		tech_service2.ModelApplePC,
		tech_service2.ModelLenovoPC,
		tech_service2.ModelHpPC,
		tech_service2.ModelAsusPC,
		tech_service2.ModelAcerPC,
		tech_service2.ModelToshibaPC,
		tech_service2.ModelDellPC,
		tech_service2.ModelCasperPC,
		tech_service2.ModelHuaweiPC,
		tech_service2.ModelMSIPC,
		tech_service2.ModelMicrosoftPC,
		tech_service2.ModelMonsterPC,
		tech_service2.ModelAsusRogPC,
		tech_service2.ModelGpad7,
		tech_service2.ModelGpad8,
		tech_service2.ModelLGOther,
		tech_service2.ModelGMTAB,
		tech_service2.ModelGMTAB4,
		tech_service2.ModelGMTAB5,
		tech_service2.ModelGMTAB10,
		tech_service2.ModelGMTAB8,
		tech_service2.ModelIphone11,
		tech_service2.ModelIphone12,
		tech_service2.ModelIphone12Pro,
		tech_service2.ModelSamsungGalaxyS7,
		tech_service2.ModelSamsungGalaxyS9,
		tech_service2.ModelLenovoM7,
		tech_service2.ModelLenovoM8,
		tech_service2.ModeliPad6,
		tech_service2.ModeliPad9,
	}

	for _, model := range models {
		if err := m.db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}
