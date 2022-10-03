package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
	"time"
)

type expertiseServiceStore struct {
	db *gorm.DB
}

type ExpertiseServiceStore interface {
	Create(model *db.ExpertiseService) error
	Update(model *db.ExpertiseService) error
	Delete(model *db.ExpertiseService) error
	FindAll() ([]db.ExpertiseService, error)
	FindByID(id int) (db.ExpertiseService, error)
	FindBy(column string, value interface{}) ([]db.ExpertiseService, error)
	FindByModelId(modelId int) ([]db.ExpertiseService, error)
	FindByBrandIdDeviceTypeId(brandId int, deviceTypeId int) ([]db.ExpertiseService, error)
	Search(query string) ([]db.ExpertiseService, error)
	Seed() error
}

func NewExpertiseServiceStore(db *gorm.DB) ExpertiseServiceStore {
	return &expertiseServiceStore{db: db}
}

func (t *expertiseServiceStore) Create(model *db.ExpertiseService) error {
	return t.db.Create(model).Error
}

func (t *expertiseServiceStore) Update(model *db.ExpertiseService) error {
	return t.db.Save(model).Error
}

func (t *expertiseServiceStore) Delete(model *db.ExpertiseService) error {
	return t.db.Delete(model).Error
}

func (t *expertiseServiceStore) FindAll() ([]db.ExpertiseService, error) {
	var models []db.ExpertiseService
	err := t.db.Find(&models).Error
	return models, err
}

func (t *expertiseServiceStore) FindByID(id int) (db.ExpertiseService, error) {
	var model db.ExpertiseService
	err := t.db.First(&model, id).Error
	return model, err
}

func (t *expertiseServiceStore) FindBy(column string, value interface{}) ([]db.ExpertiseService, error) {
	var models []db.ExpertiseService
	err := t.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

// TODO : day alani su an icin time.Monday (1) olarak sabit gonderilemektedir. Yeterli data eklendiginde time.Now().Weekday()
func (t *expertiseServiceStore) FindByModelId(modelId int) ([]db.ExpertiseService, error) {
	var expertiseServices []db.ExpertiseService
	err := t.db.Joins("INNER JOIN expertise_services_models on expertise_services.id = expertise_services_models.expertise_service_id").Where("expertise_services_models.model_id = ?", modelId).Preload("ExpertiseServiceShifts", "day = ?", time.Monday).Preload("ExpertiseServiceReservations").Find(&expertiseServices).Error
	return expertiseServices, err
}

func (t *expertiseServiceStore) FindByBrandIdDeviceTypeId(brandId int, deviceTypeId int) ([]db.ExpertiseService, error) {
	var expertiseServices []db.ExpertiseService
	err := t.db.Joins("INNER JOIN expertise_services_brands on expertise_services.id = expertise_services_brands.expertise_service_id").Where("expertise_services_brands.brand_id = ?", brandId).
		Joins("INNER JOIN expertise_services_device_types on expertise_services.id = expertise_services_device_types.expertise_service_id").Where("expertise_services_device_types.device_type_id = ?", deviceTypeId).
		Preload("ExpertiseServiceShifts", "day = ?", time.Now().Weekday()).
		Preload("Reservations").
		Find(&expertiseServices).Error
	return expertiseServices, err
}

/*
func (t *expertiseServiceStore) FindShifts(expertiseServiceId uint) ([]db.ExpertiseService, error) {
	var expertiseServices []db.ExpertiseService
	err := t.db.Model(&expertiseServices).Joins("INNER JOIN expertise_services_models on model_id = expertise_services_models.model_id").Where("expertise_services_models.model_id = ?", modelId).Error
	return expertiseServices, err
}

func (t *expertiseServiceStore) FindReservations(expertiseServiceId uint, dateTime time.Time) ([]db.ExpertiseService, error) {
	var expertiseServices []db.ExpertiseService
	err := t.db.Model(&expertiseServices).Joins("INNER JOIN expertise_services_models on model_id = expertise_services_models.model_id").Where("expertise_services_models.model_id = ?", modelId).Error
	return expertiseServices, err
}*/

func (t *expertiseServiceStore) Search(query string) ([]db.ExpertiseService, error) {
	var models []db.ExpertiseService
	err := t.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (t *expertiseServiceStore) Seed() error {
	t.db.Create(&seed_data.ExpertiseService1)
	t.db.Create(&seed_data.ExpertiseService2)
	t.db.Create(&seed_data.ExpertiseService3)

	return nil
}
