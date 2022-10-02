package web

import "time"

type ExpertiseServiceRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

type ExpertiseServiceUpdateRequest struct {
	ServiceName string `json:"service_name"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	About       string `json:"about"`
	Address     string `json:"address"`
}

type ExpertiseServiceResponse struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

type ExpertiseServiceSearchRequest struct {
	ModelId   int    `json:"model_id"`
	FixTypeId string `json:"fix_type"`
}

type ExpertiseServiceSearchResponse struct {
	Id                           int                           `json:"id"`
	Name                         string                        `json:"name"`
	Address                      string                        `json:"address"`
	Price                        uint64                        `json:"price"`
	ExpertiseServiceShift        ExpertiseServiceShift         `json:"expertise_service_shift"`
	ExpertiseServiceReservations []ExpertiseServiceReservation `json:"expertise_service_reservation"`
}

type ExpertiseServiceShift struct {
	Day          time.Weekday `json:"day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}

type ExpertiseServiceReservation struct {
	Day          time.Weekday `json:"day"`
	DateOfDay    time.Time    `json:"date_of_day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}
