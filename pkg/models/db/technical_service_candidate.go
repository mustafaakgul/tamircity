package db

import "gorm.io/gorm"

type TechnicalServiceCandidate struct {
	gorm.Model
	ServiceName         string `gorm:"type:varchar(50);not null"`
	BusinessType        string `gorm:"type:varchar(50);not null"`
	Address             string `gorm:"type:varchar(50);not null"`
	NumberofBranches    int    `gorm:"type:int;not null"`
	NumberofTechnicians int    `gorm:"type:int;not null"`
	Name                string `gorm:"type:varchar(50);not null"`
	Surname             string `gorm:"type:varchar(50);not null"`
	Email               string `gorm:"type:varchar(50);not null"`
	PhoneNumber         string `gorm:"type:varchar(15);not null"`
	IsActive            bool   `gorm:"type:boolean;default:true;not null"`
	Status              TechnicalServiceCandidateStatus
}

type TechnicalServiceCandidateStatus string

const (
	TechnicalServiceCandidateStatus_PENDING   TechnicalServiceCandidateStatus = "Pending"
	TechnicalServiceCandidateStatus_CANCELLED TechnicalServiceCandidateStatus = "Cancelled"
	TechnicalServiceCandidateStatus_APPROVED  TechnicalServiceCandidateStatus = "Approved"
)
