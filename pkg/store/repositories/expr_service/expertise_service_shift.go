package expr_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/expr_service"
	"gorm.io/gorm"
)

type expertiseServiceShiftStore struct {
	db *gorm.DB
}

type ExpertiseServiceShiftStore interface {
	CreateOrUpdate(expertiseServiceId int, expertiseServiceShift expr_service.ExpertiseServiceShift) error
	FindByExpertiseServiceId(expertiseServiceId int) ([]expr_service.ExpertiseServiceShift, error)
}

func NewExpertiseServiceShiftStore(db *gorm.DB) ExpertiseServiceShiftStore {
	return &expertiseServiceShiftStore{db: db}
}

func (t *expertiseServiceShiftStore) CreateOrUpdate(expertiseServiceId int, expertiseServiceShift expr_service.ExpertiseServiceShift) error {
	tx := t.db.Where(expr_service.ExpertiseServiceShift{ExpertiseServiceId: uint(expertiseServiceId), Day: expertiseServiceShift.Day, Status: true}).Assign(expr_service.ExpertiseServiceShift{StartOfShift: expertiseServiceShift.StartOfShift, EndOfShift: expertiseServiceShift.EndOfShift, Status: true}).FirstOrCreate(&expertiseServiceShift)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *expertiseServiceShiftStore) FindByExpertiseServiceId(expertiseServiceId int) (expertiseServiceShifts []expr_service.ExpertiseServiceShift, err error) {
	if err := t.db.Model(&expertiseServiceShifts).Where("expertise_service_id = ?", expertiseServiceId).Find(&expertiseServiceShifts).Error; err != nil {
		return nil, err
	}
	return expertiseServiceShifts, nil
}
