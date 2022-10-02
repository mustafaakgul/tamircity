package expr_service

import "gorm.io/gorm"

type DeviceType struct {
	gorm.Model
	Name              string              `gorm:"type:varchar(255);not null"`
	ShortDescription  string              `gorm:"type:varchar(255);"`
	IsActive          bool                `gorm:"type:boolean;default:true;not null"`
	ExpertiseServices []*ExpertiseService `gorm:"many2many:expertise_services_device_types;"`
	Brands            []*Brand            `gorm:"many2many:device_types_brands;"`
	Models            []*Model            //`gorm:"many2many:models_device_types;"`
	Reservation       []*Reservation
}
