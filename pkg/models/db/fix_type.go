package db

import "gorm.io/gorm"

type FixType struct {
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	Is_Active   bool   `json:"is_active" gorm:"type:boolean;default:true;not null"`
}
