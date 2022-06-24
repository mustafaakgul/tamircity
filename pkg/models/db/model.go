package db

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Name             string `json:"name" gorm:"type:varchar(255);not null"`
	ShortDescription string `json:"short_description" gorm:"type:varchar(255);"`
	BrandId          int    `json:"brand_id"`
	Brand            Brand  `json:"brand" gorm:"foreignkey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsActive         bool   `json:"is_active" gorm:"type:boolean;default:true;not null"`
}
