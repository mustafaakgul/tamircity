package db

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Is_Active bool   `json:"is_active" gorm:"type:boolean;default:true;not null"`
}
