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
	UpdateReservationStatus(reservationId int, status db.ReservationStatus) error
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

func (r *reservationStore) UpdateReservationStatus(reservationId int, status db.ReservationStatus) error {
	var reservation db.Reservation
	if err := r.db.First(&reservation, reservationId).Error; err != nil {
		return err
	}
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(&db.Reservation{}).Where("id = ?", reservationId).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}
	if status == 2 { // TO DO : Enum
		if err := tx.Create(&db.TechnicalServiceReservation{TechnicalServiceId: reservation.ID, Day: reservation.ReservationDate.Weekday(), DateofDay: reservation.ReservationDate, StartOfShift: reservation.ReservationDate, EndOfShift: reservation.ReservationDate}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
