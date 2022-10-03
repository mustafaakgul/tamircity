package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
	"time"
)

type technicalServiceStore struct {
	db *gorm.DB
}

type TechnicalServiceStore interface {
	Create(model *tech_service.TechnicalService) error
	Update(model *tech_service.TechnicalService) error
	Delete(model *tech_service.TechnicalService) error
	FindAll() ([]tech_service.TechnicalService, error)
	FindByID(id int) (tech_service.TechnicalService, error)
	FindBy(column string, value interface{}) ([]tech_service.TechnicalService, error)
	FindByModelId(modelId int) ([]tech_service.TechnicalService, error)
	Search(query string) ([]tech_service.TechnicalService, error)
	Seed() error
}

func NewTechnicalServiceStore(db *gorm.DB) TechnicalServiceStore {
	return &technicalServiceStore{db: db}
}

func (t *technicalServiceStore) Create(model *tech_service.TechnicalService) error {
	return t.db.Create(model).Error
}

func (t *technicalServiceStore) Update(model *tech_service.TechnicalService) error {
	return t.db.Save(model).Error
}

func (t *technicalServiceStore) Delete(model *tech_service.TechnicalService) error {
	return t.db.Delete(model).Error
}

func (t *technicalServiceStore) FindAll() ([]tech_service.TechnicalService, error) {
	var models []tech_service.TechnicalService
	err := t.db.Find(&models).Error
	return models, err
}

func (t *technicalServiceStore) FindByID(id int) (tech_service.TechnicalService, error) {
	var model tech_service.TechnicalService
	err := t.db.First(&model, id).Error
	return model, err
}

func (t *technicalServiceStore) FindBy(column string, value interface{}) ([]tech_service.TechnicalService, error) {
	var models []tech_service.TechnicalService
	err := t.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

// TODO : day alani su an icin time.Monday (1) olarak sabit gonderilemektedir. Yeterli data eklendiginde time.Now().Weekday()
func (t *technicalServiceStore) FindByModelId(modelId int) ([]tech_service.TechnicalService, error) {
	var technicalServices []tech_service.TechnicalService
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

func (t *technicalServiceStore) Search(query string) ([]tech_service.TechnicalService, error) {
	var models []tech_service.TechnicalService
	err := t.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (t *technicalServiceStore) Seed() error {
	t.db.Create(&tech_service2.TechnicalService1)
	t.db.Create(&tech_service2.TechnicalService2)
	t.db.Create(&tech_service2.TechnicalService3)

	return nil
}
