package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type paymentInfoStore struct {
	db *gorm.DB
}

type PaymentInfoStore interface {
	Create(comment *db.PaymentInfo) error
	FindAllByExpertiseServiceId(expertiseServiceId int) ([]db.PaymentInfo, error)
	FindByID(id int) (db.PaymentInfo, error)
}

func NewPaymentInfoStore(db *gorm.DB) PaymentInfoStore {
	return &paymentInfoStore{db: db}
}

func (c *paymentInfoStore) Create(paymentInfo *db.PaymentInfo) error {
	return c.db.Create(paymentInfo).Error
}

func (c *paymentInfoStore) FindAllByExpertiseServiceId(expertiseServiceId int) ([]db.PaymentInfo, error) {
	var paymentInfos []db.PaymentInfo
	err := c.db.Model(&paymentInfos).Joins("inner join reservations on payment_infos.reservation_id = reservations.id").Where("reservations.expertise_service_id = ?", expertiseServiceId).Find(&paymentInfos).Error
	return paymentInfos, err
}

func (c *paymentInfoStore) FindByID(id int) (db.PaymentInfo, error) {
	var paymentInfo db.PaymentInfo
	err := c.db.First(&paymentInfo, id).Error
	return paymentInfo, err
}
