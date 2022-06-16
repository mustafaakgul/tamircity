package db

import "gorm.io/gorm"

type DeviceType struct {
	gorm.Model
	Name              string `json:"name" gorm:"type:varchar(255);not null"`
	Short_Description string `json:"short_description" gorm:"type:varchar(255);"`
	Is_Active         bool   `json:"is_active" gorm:"type:boolean;default:true;not null"`
}
