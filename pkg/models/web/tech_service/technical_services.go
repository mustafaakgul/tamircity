package tech_service

import "time"

type TechnicalServiceRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

type TechnicalServiceUpdateRequest struct {
	ServiceName string `json:"service_name"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	About       string `json:"about"`
	Address     string `json:"address"`
}

type TechnicalServiceResponse struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

type TechnicalServiceSearchRequest struct {
	ModelId   int    `json:"model_id"`
	FixTypeId string `json:"fix_type"`
}

type TechnicalServiceSearchResponse struct {
	Id                           int                           `json:"id"`
	Name                         string                        `json:"name"`
	Address                      string                        `json:"address"`
	Price                        uint64                        `json:"price"`
	TechnicalServiceShift        TechnicalServiceShift         `json:"technical_service_shift"`
	TechnicalServiceReservations []TechnicalServiceReservation `json:"technical_service_reservation"`
}

type TechnicalServiceShift struct {
	Day          time.Weekday `json:"day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}

type TechnicalServiceReservation struct {
	Day          time.Weekday `json:"day"`
	DateOfDay    time.Time    `json:"date_of_day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}
