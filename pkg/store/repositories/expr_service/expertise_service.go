package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
	"time"
)

type expertiseServiceStore struct {
	db *gorm.DB
}

type ExpertiseServiceStore interface {
	Create(model *expr_service.ExpertiseService) error
	Update(model *expr_service.ExpertiseService) error
	Delete(model *expr_service.ExpertiseService) error
	FindAll() ([]expr_service.ExpertiseService, error)
	FindByID(id int) (expr_service.ExpertiseService, error)
	FindBy(column string, value interface{}) ([]expr_service.ExpertiseService, error)
	FindByModelId(modelId int) ([]expr_service.ExpertiseService, error)
	Search(query string) ([]expr_service.ExpertiseService, error)
	Seed() error
}

func NewExpertiseServiceStore(db *gorm.DB) ExpertiseServiceStore {
	return &expertiseServiceStore{db: db}
}

func (t *expertiseServiceStore) Create(model *expr_service.ExpertiseService) error {
	return t.db.Create(model).Error
}

func (t *expertiseServiceStore) Update(model *expr_service.ExpertiseService) error {
	return t.db.Save(model).Error
}

func (t *expertiseServiceStore) Delete(model *expr_service.ExpertiseService) error {
	return t.db.Delete(model).Error
}

func (t *expertiseServiceStore) FindAll() ([]expr_service.ExpertiseService, error) {
	var models []expr_service.ExpertiseService
	err := t.db.Find(&models).Error
	return models, err
}

func (t *expertiseServiceStore) FindByID(id int) (expr_service.ExpertiseService, error) {
	var model expr_service.ExpertiseService
	err := t.db.First(&model, id).Error
	return model, err
}

func (t *expertiseServiceStore) FindBy(column string, value interface{}) ([]expr_service.ExpertiseService, error) {
	var models []expr_service.ExpertiseService
	err := t.db.Where(column+" = ?", value).Find(&models).Error
	return models, err
}

// TODO : day alani su an icin time.Monday (1) olarak sabit gonderilemektedir. Yeterli data eklendiginde time.Now().Weekday()
func (t *expertiseServiceStore) FindByModelId(modelId int) ([]expr_service.ExpertiseService, error) {
	var expertiseServices []expr_service.ExpertiseService
	err := t.db.Joins("INNER JOIN expertise_services_models on expertise_services.id = expertise_services_models.expertise_service_id").Where("expertise_services_models.model_id = ?", modelId).Preload("ExpertiseServiceShifts", "day = ?", time.Monday).Preload("ExpertiseServiceReservations").Find(&expertiseServices).Error
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

func (t *expertiseServiceStore) Search(query string) ([]expr_service.ExpertiseService, error) {
	var models []expr_service.ExpertiseService
	err := t.db.Where("name LIKE ?", "%"+query+"%").Find(&models).Error
	return models, err
}

func (t *expertiseServiceStore) Seed() error {
	return nil
}
