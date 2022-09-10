package db

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name             string      `gorm:"type:varchar(255);not null"`
	ShortDescription string      `gorm:"type:varchar(255);"`
	IsActive         bool        `gorm:"type:boolean;default:true;not null"`
	BrandID          *uint       `gorm:"not null"` //`gorm:"default:null;"`
	Brand            *Brand      //`gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,default:null;"`
	DeviceTypeId     *uint       `gorm:"default:null;"`
	DeviceType       *DeviceType //`gorm:"foreignkey:DeviceTypeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,default:null;"`
	//Brands           []*Brand            `gorm:"many2many:models_brands;"`
	//DeviceTypes      []*DeviceType       `gorm:"many2many:models_device_types;"`
	TechnicalService []*TechnicalService `gorm:"many2many:technical_services_models;"`
	Reservations     []*Reservation
}
