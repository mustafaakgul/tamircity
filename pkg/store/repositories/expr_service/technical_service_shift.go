package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"gorm.io/gorm"
)

type technicalServiceShiftStore struct {
	db *gorm.DB
}

type TechnicalServiceShiftStore interface {
	CreateOrUpdate(technicalServiceId int, technicalServiceShift tech_service.TechnicalServiceShift) error
	FindByTechnicalServiceId(technicalServiceId int) ([]tech_service.TechnicalServiceShift, error)
}

func NewTechnicalServiceShiftStore(db *gorm.DB) TechnicalServiceShiftStore {
	return &technicalServiceShiftStore{db: db}
}

func (t *technicalServiceShiftStore) CreateOrUpdate(technicalServiceId int, technicalServiceShift tech_service.TechnicalServiceShift) error {
	tx := t.db.Where(tech_service.TechnicalServiceShift{TechnicalServiceId: uint(technicalServiceId), Day: technicalServiceShift.Day, Status: true}).Assign(tech_service.TechnicalServiceShift{StartOfShift: technicalServiceShift.StartOfShift, EndOfShift: technicalServiceShift.EndOfShift, Status: true}).FirstOrCreate(&technicalServiceShift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *technicalServiceShiftStore) FindByTechnicalServiceId(technicalServiceId int) (technicalServiceShifts []tech_service.TechnicalServiceShift, err error) {
	if err := t.db.Model(&technicalServiceShifts).Where("technical_service_id = ?", technicalServiceId).Find(&technicalServiceShifts).Error; err != nil {
		return nil, err
	}
	return technicalServiceShifts, nil
}
