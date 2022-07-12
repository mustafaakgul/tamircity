package db

import "gorm.io/gorm"

type DeviceType struct {
	gorm.Model
	Name              string              `json:"name" gorm:"type:varchar(255);not null"`
	ShortDescription  string              `json:"short_description" gorm:"type:varchar(255);"`
	IsActive          bool                `json:"is_active" gorm:"type:boolean;default:true;not null"`
	TechnicalServices []*TechnicalService `gorm:"many2many:technical_services_device_types;"`
}
