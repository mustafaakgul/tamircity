package db

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	DeviceTypeId       int `gorm:"not null"`
	BrandId            int `gorm:"not null"`
	ModelId            int `gorm:"not null"`
	FixTypeId          int `gorm:"not null"`
	ServiceTypeId      int `gorm:"not null"`
	ExtraServiceId     int `gorm:"not null"`
	TechnicalServiceId int
	ReservationDate    time.Time
	Price              int
	Status             ReservationStatus
	FullName           string `gorm:"not null"`
	Email              string `gorm:"not null"`
	PhoneNumber        string `gorm:"not null"`
	SecondPhoneNumber  string `gorm:"not null"`
	Description        string
}

type ReservationStatus int

const (
	Pending  ReservationStatus = 0 // Beklemede
	Denied                     = 1 // Reddedildi
	Approved                   = 2 // Onaylandı
)
