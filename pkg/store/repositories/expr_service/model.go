package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type modelStore struct {
	db *gorm.DB
}

type ModelStore interface {
	Migration()
	Create(model *expr_service.Model) error
	Update(model *expr_service.Model) error
	Delete(model *expr_service.Model) error
	FindAll() ([]expr_service.Model, error)
	FindByID(id int) (expr_service.Model, error)
	FindBy(column string, value interface{}) ([]expr_service.Model, error)
	Where(brandId int, deviceTypeId int) ([]expr_service.Model, error)
	Search(query string) ([]expr_service.Model, error)
	Seed() error
}

func NewModelStore(db *gorm.DB) ModelStore {
	return &modelStore{db: db}
}

func (m *modelStore) Migration() {
	m.db.AutoMigrate(&expr_service.Model{})
}

func (m *modelStore) Create(model *expr_service.Model) error {
	return m.db.Create(model).Error
}

func (m *modelStore) Update(model *expr_service.Model) error {
	return m.db.Save(model).Error
}

func (m *modelStore) Delete(model *expr_service.Model) error {
	return m.db.Delete(model).Error
}

func (m *modelStore) FindAll() ([]expr_service.Model, error) {
	var models []expr_service.Model
	err := m.db.Find(&models).Error
	return models, err
}

func (m *modelStore) FindByID(id int) (expr_service.Model, error) {
	var model expr_service.Model
	err := m.db.First(&model, id).Error
	return model, err
}

func (m *modelStore) FindBy(column string, value interface{}) ([]expr_service.Model, error) {
	var models []expr_service.Model
	err := m.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

func (m *modelStore) Where(brandId int, deviceTypeId int) ([]expr_service.Model, error) {
	var models []expr_service.Model
	/*err := m.db.Model(&models).
	Joins("inner join models_brands on models.id = models_brands.model_id").
	Where("models_brands.brand_id = ?", brandId).
	Joins("inner join models_device_types on models.id = models_device_types.model_id").
	Where("models_device_types.device_type_id = ?", deviceTypeId).
	Find(&models).Error*/
	err := m.db.Where("brand_id = ? AND device_type_id = ?", brandId, deviceTypeId).Find(&models).Error
	return models, err
}

func (m *modelStore) Search(query string) ([]expr_service.Model, error) {
	var models []expr_service.Model
	err := m.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (m *modelStore) Seed() error {
	return nil
}
