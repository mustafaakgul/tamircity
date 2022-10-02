package db

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name              string              `gorm:"type:varchar(255);not null"`
	IsActive          bool                `gorm:"type:boolean;default:true;not null"`
	ExpertiseServices []*ExpertiseService `gorm:"many2many:expertise_services_brands;"`
	DeviceTypes       []*DeviceType       `gorm:"many2many:device_types_brands;"`
	Models            []*Model            //`gorm:"many2many:models_brands;"`
	Reservations      []*Reservation
}
