package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type expertiseServiceShiftStore struct {
	db *gorm.DB
}

type ExpertiseServiceShiftStore interface {
	CreateOrUpdate(expertiseServiceId int, expertiseServiceShift db.ExpertiseServiceShift) error
	FindByExpertiseServiceId(expertiseServiceId int) ([]db.ExpertiseServiceShift, error)
}

func NewExpertiseServiceShiftStore(db *gorm.DB) ExpertiseServiceShiftStore {
	return &expertiseServiceShiftStore{db: db}
}

func (t *expertiseServiceShiftStore) CreateOrUpdate(expertiseServiceId int, expertiseServiceShift db.ExpertiseServiceShift) error {
	tx := t.db.Where(db.ExpertiseServiceShift{ExpertiseServiceId: uint(expertiseServiceId), Day: expertiseServiceShift.Day, Status: true}).Assign(db.ExpertiseServiceShift{StartOfShift: expertiseServiceShift.StartOfShift, EndOfShift: expertiseServiceShift.EndOfShift, Status: true}).FirstOrCreate(&expertiseServiceShift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *expertiseServiceShiftStore) FindByExpertiseServiceId(expertiseServiceId int) (expertiseServiceShifts []db.ExpertiseServiceShift, err error) {
	if err := t.db.Model(&expertiseServiceShifts).Where("expertise_service_id = ?", expertiseServiceId).Find(&expertiseServiceShifts).Error; err != nil {
		return nil, err
	}
	return expertiseServiceShifts, nil
}
