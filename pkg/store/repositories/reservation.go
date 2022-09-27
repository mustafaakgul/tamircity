package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/store/seed_data"
	"gorm.io/gorm"
	"time"
)

type reservationStore struct {
	db *gorm.DB
}

type ReservationStore interface {
	Create(reservation *db.Reservation) error
	GetPendingListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error)
	GetCompletedListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error)
	GetCancelledListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error)
	GetApprovedListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error)
	GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) ([]db.Reservation, error)
	GetReservationCountWithStatus(technicalServiceId int, status db.ReservationStatus) (count int64, err error)
	UpdateReservationStatus(reservationId int, status db.ReservationStatus) error
	ChangeOperationStatus(reservationId int, operationStatus db.OperationStatus) error
	Seed() error
}

func NewReservationStore(db *gorm.DB) ReservationStore {
	return &reservationStore{db: db}
}

func (r *reservationStore) Create(reservation *db.Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *reservationStore) GetPendingListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, db.Pending).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCompletedListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, db.Completed).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, db.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByTechnicalServiceIdAndDatetime(technicalServiceId int, reservationDate time.Time) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND reservation_date > ? AND reservation_date < ?  AND status = ?  ", technicalServiceId, reservationDate, reservationDate.Add(time.Hour*24), db.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCancelledListByTechnicalServiceId(technicalServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("technical_service_id  = ? AND status = ? ", technicalServiceId, db.Cancelled).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("TechnicalService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) UpdateReservationStatus(reservationId int, status db.ReservationStatus) error {

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

	if err := r.db.Model(&db.Reservation{}).Where("id = ?", reservationId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *reservationStore) ChangeOperationStatus(reservationId int, operationStatus db.OperationStatus) error {
	return r.db.Model(&db.Reservation{}).Where("id = ?", reservationId).Update("operational_status", operationStatus).Error
}

func (r *reservationStore) GetReservationCountWithStatus(technicalServiceId int, status db.ReservationStatus) (count int64, err error) {
	if err := r.db.Model(&db.Reservation{}).Where("technical_service_id = ? AND status = ?", technicalServiceId, status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *reservationStore) Seed() error {
	r.db.Create(&seed_data.Reservation1)
	r.db.Create(&seed_data.Reservation2)
	r.db.Create(&seed_data.Reservation3)
	r.db.Create(&seed_data.Reservation4)
	r.db.Create(&seed_data.Reservation5)
	r.db.Create(&seed_data.Reservation6)
	return nil
}
