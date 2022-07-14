package db

import "gorm.io/gorm"

type DeviceType struct {
	gorm.Model
	Name              string              `gorm:"type:varchar(255);not null"`
	ShortDescription  string              `gorm:"type:varchar(255);"`
	IsActive          bool                `gorm:"type:boolean;default:true;not null"`
	TechnicalServices []*TechnicalService `gorm:"many2many:technical_services_device_types;"`
	Brands            []*Brand            `gorm:"many2many:device_types_brands;"`
	FixTypes          []*FixType          `gorm:"many2many:device_types_fix_types;"`
	Models            []*Model
}
