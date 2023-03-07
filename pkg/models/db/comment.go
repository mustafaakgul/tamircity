package db

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	CommentOwner       string `gorm:"type:varchar(100);not null"`
	Comment            string `gorm:"type:varchar(200);not null"`
	CommentDate        time.Time
	Rate               int `gorm:"type:int;not null"`
	ExpertiseServiceID int
	ExpertiseService   *ExpertiseService `gorm:"foreignkey:ExpertiseServiceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
