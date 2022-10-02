package tech_service

import (
	"gorm.io/gorm"
	"time"
)

type TechnicalServiceReservation struct {
	gorm.Model
	TechnicalServiceId uint
	TechnicalService   TechnicalService `gorm:"foreignkey:TechnicalServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Day                time.Weekday     // Enum (Pazartesi, Salı, Çarşamba vs.)
	DateofDay          time.Time
	StartOfShift       int
	EndOfShift         int
}
