package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
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

func (t *technicalServiceStore) FindByModelId(modelId int) ([]db.TechnicalService, error) {
	var technicalServices []db.TechnicalService
	err := t.db.Joins("INNER JOIN technical_services_models on technical_services.id = technical_services_models.technical_service_id").Where("technical_services_models.model_id = ?", modelId).Preload("TechnicalServiceShifts", "day = ?", time.Now().Day()).Preload("TechnicalServiceReservations").Find(&technicalServices).Error
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

var deviceTypePc = &db.DeviceType{
	Name:             "Personel Computer",
	ShortDescription: "PC",
	IsActive:         true,
}

var deviceTypePhone = &db.DeviceType{
	Name:             "Phone",
	ShortDescription: "Phone",
	IsActive:         true,
}

var deviceTypeTablet = &db.DeviceType{
	Name:             "Tablet",
	ShortDescription: "Tablet",
	IsActive:         true,
}

//var deviceTypePc = []*db.DeviceType{
//	{
//		Name:             "Personel Computer",
//		ShortDescription: "PC",
//		IsActive:         true,
//	},
//}

//var deviceTypePhone = []*db.DeviceType{
//	{
//		Name:             "Phone",
//		ShortDescription: "Phone",
//		IsActive:         true,
//	},
//}
//var deviceTypePhone = db.DeviceType{
//	Name:             "Phone",
//	ShortDescription: "Phone",
//	IsActive:         true,
//}
//
//var deviceTypeTablet = []*db.DeviceType{
//	{
//		Name:             "Tablet",
//		ShortDescription: "Tablet",
//		IsActive:         true,
//	},
//}

//var deviceTypePhone = []*db.DeviceType{
//	{
//		Name:             "Phone",
//		ShortDescription: "Phone",
//		IsActive:         true,
//	},
//}
//
//var deviceTypeTablet = []*db.DeviceType{
//	{
//		Name:             "Tablet",
//		ShortDescription: "Tablet",
//		IsActive:         true,
//	},
//}

func (t *technicalServiceStore) Seed() error {

	technicalServices := []db.TechnicalService{
		{
			ServiceName:    "service1",
			IdentityNumber: "123456789",
			PhoneNumber:    "123456789",
			Email:          "email1@email.com",
			Iban:           "123456789",
			IsActive:       true,
			DeviceTypes:    []*db.DeviceType{deviceTypePc},
		},
		{
			ServiceName:    "service2",
			IdentityNumber: "123456789",
			PhoneNumber:    "123456789",
			Email:          "email2@email.com",
			Iban:           "123456789",
			IsActive:       true,
			DeviceTypes:    []*db.DeviceType{deviceTypePhone},
		},
		{
			ServiceName:    "service3",
			IdentityNumber: "123456789",
			PhoneNumber:    "123456789",
			Email:          "email3@email.com",
			Iban:           "123456789",
			IsActive:       true,
			DeviceTypes:    []*db.DeviceType{deviceTypeTablet},
		},
	}

	for _, technicalService := range technicalServices {
		if err := t.db.Create(&technicalService).Error; err != nil {
			return err
		}
	}

	return nil
}
