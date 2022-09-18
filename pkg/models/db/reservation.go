package db

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	DeviceTypeId       int          `gorm:"not null"`
	DeviceType         DeviceType   `gorm:"foreignkey:DeviceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BrandId            int          `gorm:"not null"`
	Brand              Brand        `gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ModelId            int          `gorm:"not null"`
	ModelEntity        Model        `gorm:"foreignkey:ModelId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FixTypeId          int          `gorm:"not null"`
	FixType            FixType      `gorm:"foreignkey:FixTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ServiceTypeId      int          `gorm:"not null"`
	ServiceType        ServiceType  `gorm:"foreignkey:ServiceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExtraServiceId     int          `gorm:"not null"`
	ExtraService       ExtraService `gorm:"foreignkey:ExtraServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TechnicalServiceId int
	TechnicalService   TechnicalService `gorm:"foreignkey:TechnicalServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReservationDate    time.Time
	StartOfHour        int
	EndOfHour          int
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
	Pending   ReservationStatus = 0 // Beklemede
	Cancelled                   = 1 // İptal Edildi
	Approved                    = 2 // Onaylandı
	Completed                   = 3 // Tamamlandı
)
