package db

import (
	"gorm.io/gorm"
	"time"
)

type ExpertiseServiceShift struct {
	gorm.Model
	ExpertiseServiceId uint
	ExpertiseService   ExpertiseService `gorm:"foreignkey:ExpertiseServiceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Day                time.Weekday     // (Pazartesi, Salı, Çarşamba vs.)
	StartOfShift       int
	EndOfShift         int
	Status             bool
}
