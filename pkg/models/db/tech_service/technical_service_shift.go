package tech_service

import (
	"gorm.io/gorm"
	"time"
)

type TechnicalServiceShift struct {
	gorm.Model
	TechnicalServiceId uint
	TechnicalService   TechnicalService `gorm:"foreignkey:TechnicalServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Day                time.Weekday     // (Pazartesi, Salı, Çarşamba vs.)
	StartOfShift       int
	EndOfShift         int
	Status             bool
}
