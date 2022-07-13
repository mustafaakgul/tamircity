package db

import "gorm.io/gorm"

type Newsletter struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null"`
}
