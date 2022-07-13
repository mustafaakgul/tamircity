package db

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FullName    string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(50);not null"`
	PhoneNumber string `gorm:"type:varchar(15);not null"`
}
