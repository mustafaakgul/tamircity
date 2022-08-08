package db

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name             string `gorm:"type:varchar(255);not null"`
	ShortDescription string `gorm:"type:varchar(255);"`
	BrandId          int
	Brand            Brand `gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeviceTypeId     int
	DeviceType       DeviceType          `gorm:"foreignkey:DeviceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsActive         bool                `gorm:"type:boolean;default:true;not null"`
	TechnicalService []*TechnicalService `gorm:"many2many:technical_services_models;"`
	Reservations     []*Reservation
}
