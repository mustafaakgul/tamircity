package db

import "gorm.io/gorm"

type TechnicalService struct {
	gorm.Model
	ServiceName    string        `gorm:"type:varchar(50);not null"`
	IdentityNumber string        `gorm:"type:varchar(13);not null"`
	PhoneNumber    string        `gorm:"type:varchar(15);not null"`
	Email          string        `gorm:"type:varchar(50);not null"`
	Iban           string        `gorm:"type:varchar(25);"`
	IsActive       bool          `gorm:"type:boolean;default:true;not null"`
	DeviceTypes    []*DeviceType `gorm:"many2many:technical_services_device_types;"`
	Brands         []*Brand      `gorm:"many2many:technical_services_brands;"`
}
