package db

import (
	"gorm.io/gorm"
)

type ServiceType struct {
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	Price       int64
	IsActive    bool `json:"is_active" gorm:"type:boolean;default:true;not null"`
}

/*
yerinde tamir
merkezde tamir
kargo
kurye
*/
