package db

import (
	"gorm.io/gorm"
)

type ServiceType struct {
	gorm.Model
	Description string `gorm:"type:varchar(255);not null"`
	Price       int64
	IsActive    bool `gorm:"type:boolean;default:true;not null"`
}

/*
yerinde tamir
merkezde tamir
kargo
kurye
*/
