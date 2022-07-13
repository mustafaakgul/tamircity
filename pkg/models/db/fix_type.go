package db

import "gorm.io/gorm"

type FixType struct {
	gorm.Model
	Description string        `gorm:"type:varchar(255);not null"`
	IsActive    bool          `gorm:"type:boolean;default:true;not null"`
	DeviceTypes []*DeviceType `gorm:"many2many:device_types_fix_types;"`
}
