package tech_service

import (
	"gorm.io/gorm"
)

type ServiceType struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	Description  string `gorm:"type:varchar(255);not null"`
	Price        int64
	IsActive     bool `gorm:"type:boolean;default:true;not null"`
	Reservations []*Reservation
}
