package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/store/seed_data/tech_service"
	"gorm.io/gorm"
	"time"
)

type reservationStore struct {
	db *gorm.DB
}

type ReservationStore interface {
	Create(reservation *tech_service.Reservation) error
	FindByID(id int) (*tech_service.Reservation, error)
	GetPendingListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error)
	GetCompletedListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error)
	GetCancelledListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error)
	GetApprovedListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error)
	GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) ([]tech_service.Reservation, error)
	GetReservationCountWithStatus(technicalServiceId int, status tech_service.ReservationStatus) (count int64, err error)
	UpdateReservationStatus(reservationId int, status tech_service.ReservationStatus) error
	ChangeOperationStatus(reservationId int, operationStatus tech_service.OperationStatus) error
	Seed() error
}

func NewReservationStore(db *gorm.DB) ReservationStore {
	return &reservationStore{db: db}
}

func (r *reservationStore) Create(reservation *tech_service.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationStore) FindByID(id int) (*tech_service.Reservation, error) {
	var reservation tech_service.Reservation
	err := r.db.First(&reservation, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservation).Error
	return &reservation, err
}

func (r *reservationStore) GetPendingListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error) {
	var reservations []tech_service.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, tech_service.Pending).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCompletedListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error) {
	var reservations []tech_service.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, tech_service.Completed).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error) {
	var reservations []tech_service.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, tech_service.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) ([]tech_service.Reservation, error) {
	var reservations []tech_service.Reservation
	err := r.db.Where("technical_service_id  = ? AND reservation_date > ? AND reservation_date < ?  AND status = ?  ", technicalServiceId, reservationDate, reservationDate.Add(time.Hour*24), tech_service.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCancelledListByTechnicalServiceId(technicalServiceId int) ([]tech_service.Reservation, error) {
	var reservations []tech_service.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, tech_service.Cancelled).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) UpdateReservationStatus(reservationId int, status tech_service.ReservationStatus) error {

	//TODO: TechnicalServiceReservation tablosu incelenip kaldirilacak.
	/*
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
			if err := tx.Create(&db.TechnicalServiceReservation{TechnicalServiceId: reservation.ID, Day: reservation.ReservationDate.Weekday(), DateofDay: reservation.ReservationDate, StartOfShift: reservation.StartOfHour, EndOfShift: reservation.EndOfHour}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	*/

	if err := r.db.Model(&tech_service.Reservation{}).Where("id = ?", reservationId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *reservationStore) ChangeOperationStatus(reservationId int, operationStatus tech_service.OperationStatus) error {
	return r.db.Model(&tech_service.Reservation{}).Where("id = ?", reservationId).Update("operational_status", operationStatus).Error
}

func (r *reservationStore) GetReservationCountWithStatus(technicalServiceId int, status tech_service.ReservationStatus) (count int64, err error) {
	if err := r.db.Model(&tech_service.Reservation{}).Where("technical_service_id = ? AND status = ?", technicalServiceId, status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *reservationStore) Seed() error {
	r.db.Create(&tech_service2.Reservation1)
	r.db.Create(&tech_service2.Reservation2)
	r.db.Create(&tech_service2.Reservation3)
	r.db.Create(&tech_service2.Reservation4)
	r.db.Create(&tech_service2.Reservation5)
	r.db.Create(&tech_service2.Reservation6)

	return nil
}
