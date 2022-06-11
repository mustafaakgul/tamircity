package db

import "gorm.io/gorm"

type FixType struct {
	gorm.Model
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	IsActive    bool   `json:"is_active" gorm:"type:boolean;not null"`
}

func (FixType) TableName() string {
	return "fixTypes"
}
