package db

import (
	"gorm.io/gorm"
	"time"
)

type PaymentInfo struct {
	gorm.Model
	PaymentDate   time.Time
	Amount        int `gorm:"type:int;not null"`
	PaymentType   PaymentType
	ReservationID int
	Reservation   *Reservation `gorm:"foreignkey:ReservationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PaymentType int

const (
	Card PaymentType = 0 // Kart
	Cash             = 1 // Nakit
)
