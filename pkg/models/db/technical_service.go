package db

import "gorm.io/gorm"

type TechnicalService struct {
	gorm.Model
	ServiceName                  string        `gorm:"type:varchar(50);not null"`
	BusinessType                 string        `gorm:"type:varchar(50);"`
	NumberofBranches             int           `gorm:"type:int;"`
	NumberofTechnicians          int           `gorm:"type:int;"`
	Name                         string        `gorm:"type:varchar(50);"`
	Surname                      string        `gorm:"type:varchar(50);"`
	IdentityNumber               string        `gorm:"type:varchar(13);not null"`
	PhoneNumber                  string        `gorm:"type:varchar(15);not null"`
	Email                        string        `gorm:"type:varchar(50);not null"`
	Iban                         string        `gorm:"type:varchar(25);"`
	IsActive                     bool          `gorm:"type:boolean;default:true;not null"`
	Address                      string        `gorm:"type:varchar(50);not null"`
	About                        string        `gorm:"type:varchar(200);not null"`
	DeviceTypes                  []*DeviceType `gorm:"many2many:technical_services_device_types;"`
	Brands                       []*Brand      `gorm:"many2many:technical_services_brands;"`
	Models                       []*Model      `gorm:"many2many:technical_services_models;"`
	TechnicalServiceShifts       []*TechnicalServiceShift
	TechnicalServiceReservations []*TechnicalServiceReservation
	Reservations                 []*Reservation
}
