package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
	"time"
)

type technicalServiceStore struct {
	db *gorm.DB
}

type TechnicalServiceStore interface {
	Create(model *db.TechnicalService) error
	Update(model *db.TechnicalService) error
	Delete(model *db.TechnicalService) error
	FindAll() ([]db.TechnicalService, error)
	FindByID(id int) (db.TechnicalService, error)
	FindBy(column string, value interface{}) ([]db.TechnicalService, error)
	FindByModelId(modelId int) ([]db.TechnicalService, error)
	Search(query string) ([]db.TechnicalService, error)
	Seed() error
}

func NewTechnicalServiceStore(db *gorm.DB) TechnicalServiceStore {
	return &technicalServiceStore{db: db}
}

func (t *technicalServiceStore) Create(model *db.TechnicalService) error {
	return t.db.Create(model).Error
}

func (t *technicalServiceStore) Update(model *db.TechnicalService) error {
	return t.db.Save(model).Error
}

func (t *technicalServiceStore) Delete(model *db.TechnicalService) error {
	return t.db.Delete(model).Error
}

func (t *technicalServiceStore) FindAll() ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Find(&models).Error
	return models, err
}

func (t *technicalServiceStore) FindByID(id int) (db.TechnicalService, error) {
	var model db.TechnicalService
	err := t.db.First(&model, id).Error
	return model, err
}

func (t *technicalServiceStore) FindBy(column string, value interface{}) ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

// TODO : day alani su an icin time.Monday (1) olarak sabit gonderilemektedir. Yeterli data eklendiginde time.Now().Weekday()
func (t *technicalServiceStore) FindByModelId(modelId int) ([]db.TechnicalService, error) {
	var technicalServices []db.TechnicalService
	err := t.db.Joins("INNER JOIN technical_services_models on technical_services.id = technical_services_models.technical_service_id").Where("technical_services_models.model_id = ?", modelId).Preload("TechnicalServiceShifts", "day = ?", time.Monday).Preload("TechnicalServiceReservations").Find(&technicalServices).Error
	return technicalServices, err
}

/*
func (t *technicalServiceStore) FindShifts(technicalServiceId uint) ([]db.TechnicalService, error) {
	var technicalServices []db.TechnicalService
	err := t.db.Model(&technicalServices).Joins("INNER JOIN technical_services_models on model_id = technical_services_models.model_id").Where("technical_services_models.model_id = ?", modelId).Error
	return technicalServices, err
}

func (t *technicalServiceStore) FindReservations(technicalServiceId uint, dateTime time.Time) ([]db.TechnicalService, error) {
	var technicalServices []db.TechnicalService
	err := t.db.Model(&technicalServices).Joins("INNER JOIN technical_services_models on model_id = technical_services_models.model_id").Where("technical_services_models.model_id = ?", modelId).Error
	return technicalServices, err
}*/

func (t *technicalServiceStore) Search(query string) ([]db.TechnicalService, error) {
	var models []db.TechnicalService
	err := t.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (t *technicalServiceStore) Seed() error {

	t.db.Create(&technicalService1)
	t.db.Create(&technicalService2)
	t.db.Create(&technicalService3)

	/*technicalServices := []*db.TechnicalService{technicalService1, technicalService2, technicalService3}

	for _, technicalService := range technicalServices {
		if err := t.db.Create(&technicalService).Error; err != nil {
			return err
		}
	}*/

	return nil
}
