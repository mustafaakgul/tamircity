package tech_service

import "gorm.io/gorm"

type FixType struct {
	gorm.Model
	Description  string `gorm:"type:varchar(255);not null"`
	IsActive     bool   `gorm:"type:boolean;default:true;not null"`
	Price        int64
	DeviceTypes  []*DeviceType `gorm:"many2many:device_types_fix_types;"`
	Reservations []*Reservation
}
