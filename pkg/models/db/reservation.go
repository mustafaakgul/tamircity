package db

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	DeviceTypeId    string `gorm:"not null"`
	BrandId         string `gorm:"not null"`
	ModelId         string `gorm:"not null"`
	FixTypeId       string `gorm:"not null"`
	ServiceTypeId   string `gorm:"not null"`
	ExtraServiceId  string `gorm:"not null"`
	ReservationDate time.Time
	FullName        string `gorm:"not null"`
	Email           string `gorm:"not null"`
	PhoneNumber     string `gorm:"not null"`
	Description     string
}
