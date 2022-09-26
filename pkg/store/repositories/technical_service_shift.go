package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type technicalServiceShiftStore struct {
	db *gorm.DB
}

type TechnicalServiceShiftStore interface {
	CreateOrUpdate(technicalServiceId int, technicalServiceShift db.TechnicalServiceShift) error
	FindByTechnicalServiceId(technicalServiceId int) ([]db.TechnicalServiceShift, error)
}

func NewTechnicalServiceShiftStore(db *gorm.DB) TechnicalServiceShiftStore {
	return &technicalServiceShiftStore{db: db}
}

func (t *technicalServiceShiftStore) CreateOrUpdate(technicalServiceId int, technicalServiceShift db.TechnicalServiceShift) error {
	tx := t.db.Where(db.TechnicalServiceShift{TechnicalServiceId: technicalServiceId, Day: technicalServiceShift.Day, Status: true}).Assign(db.TechnicalServiceShift{StartOfShift: technicalServiceShift.StartOfShift, EndOfShift: technicalServiceShift.EndOfShift, Status: true}).FirstOrCreate(&technicalServiceShift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *technicalServiceShiftStore) FindByTechnicalServiceId(technicalServiceId int) (technicalServiceShifts []db.TechnicalServiceShift, err error) {
	if err := t.db.Model(&technicalServiceShifts).Where("technical_service_id = ?", technicalServiceId).Find(&technicalServiceShifts).Error; err != nil {
		return nil, err
	}
	return technicalServiceShifts, nil
}
