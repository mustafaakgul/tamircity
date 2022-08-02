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
}

func NewReservationStore(db *gorm.DB) ReservationStore {
	return &reservationStore{db: db}
}

func (r *reservationStore) Create(reservation *db.Reservation) error {
	return r.db.Create(reservation).Error
}
