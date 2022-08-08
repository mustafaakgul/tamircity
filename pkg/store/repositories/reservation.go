package repositories

import (
	"github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
)

type reservationStore struct {
	db *gorm.DB
}

type ReservationStore interface {
	Create(reservation *db.Reservation) error
	GetPendingListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error)
}

func NewReservationStore(db *gorm.DB) ReservationStore {
	return &reservationStore{db: db}
}

func (r *reservationStore) Create(reservation *db.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationStore) GetPendingListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, 0).Preload("DeviceType").Preload("Brand").Preload("Model").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error // TO DO : 0 olan yer ReservationStatus.Pending olmalı
	return reservations, err
}
