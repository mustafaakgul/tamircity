package web

type TechnicalServiceRequest struct {
	ServiceName    string `json:"service_name" gorm:"type:varchar(50);not null"`
	IdentityNumber string `json:"identity_number" gorm:"type:varchar(13);not null"`
	PhoneNumber    string `json:"phone_number" gorm:"type:varchar(15);not null"`
	Email          string `json:"email" gorm:"type:varchar(50);not null"`
	Iban           string `json:"iban" gorm:"type:varchar(25);"`
}
