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
	FindByID(id int) (*db.Reservation, error)
	GetPendingListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error)
	GetCompletedListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error)
	GetCancelledListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error)
	GetApprovedListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error)
	GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId int, reservationDate time.Time) ([]db.Reservation, error)
	GetReservationCountWithStatus(expertiseServiceId int, status db.ReservationStatus) (count int64, err error)
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

func (r *reservationStore) FindByID(id int) (*db.Reservation, error) {
	var reservation db.Reservation
	err := r.db.First(&reservation, id).Error
	//err := r.db.Where("id = ?", id).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservation).Error
	return &reservation, err
}

func (r *reservationStore) GetPendingListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("expertise_service_id  = ? AND status = ? ", expertiseServiceId, db.Pending).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("ServiceType").Preload("ExpertiseService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCompletedListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("expertise_service_id  = ? AND status = ? ", expertiseServiceId, db.Completed).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("ServiceType").Preload("ExpertiseService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("expertise_service_id  = ? AND status = ? ", expertiseServiceId, db.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("ServiceType").Preload("ExpertiseService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetCancelledListByExpertiseServiceId(expertiseServiceId int) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("expertise_service_id  = ? AND status = ? ", expertiseServiceId, db.Cancelled).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("ServiceType").Preload("ExpertiseService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) GetApprovedListByExpertiseServiceIdAndDatetime(expertiseServiceId int, reservationDate time.Time) ([]db.Reservation, error) {
	var reservations []db.Reservation
	err := r.db.Where("expertise_service_id  = ? AND reservation_date > ? AND reservation_date < ?  AND status = ?  ", expertiseServiceId, reservationDate, reservationDate.Add(time.Hour*24), db.Approved).Preload("DeviceType").Preload("Brand").Preload("ModelEntity").Preload("FixType").Preload("ServiceType").Preload("ExtraService").Preload("ExpertiseService").Find(&reservations).Error
	return reservations, err
}

func (r *reservationStore) UpdateReservationStatus(reservationId int, status db.ReservationStatus) error {

	//TODO: ExpertiseServiceReservation tablosu incelenip kaldirilacak.
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
			if err := tx.Create(&db.ExpertiseServiceReservation{ExpertiseServiceId: reservation.ID, Day: reservation.ReservationDate.Weekday(), DateofDay: reservation.ReservationDate, StartOfShift: reservation.StartOfHour, EndOfShift: reservation.EndOfHour}).Error; err != nil {
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

func (r *reservationStore) GetReservationCountWithStatus(expertiseServiceId int, status db.ReservationStatus) (count int64, err error) {
	if err := r.db.Model(&db.Reservation{}).Where("expertise_service_id = ? AND status = ?", expertiseServiceId, status).Count(&count).Error; err != nil {
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
