package web

import "time"

type ReservationCreateRequest struct {
	DeviceTypeId       int       `json:"device_type_id"`
	BrandId            int       `json:"brand_id"`
	ModelId            int       `json:"model_id"`
	FixTypeId          int       `json:"fix_type_id"`
	ServiceTypeId      int       `json:"service_type_id"`
	ExtraServiceId     int       `json:"extra_service_id"`
	TechnicalServiceId int       `json:"technical_service_id"`
	ReservationDate    time.Time `json:"reservation_date"`
	Price              int       `json:"price"`
	FullName           string    `json:"full_name"`
	PhoneNumber        string    `json:"phone_number"`
	SecondPhoneNumber  string    `json:"second_phone_number"`
	Description        string    `json:"second_phone"`
}

type ReservationPendingResponse struct {
	ReservationId    int       `json:"reservation_id"`
	ReservationDate  time.Time `json:"reservation_date"`
	DeviceTypeName   string    `json:"device_type_name"`
	BrandName        string    `json:"brand_name"`
	ModelName        string    `json:"model_name"`
	FixTypeName      string    `json:"fix_type_name"`
	ServiceTypeName  string    `json:"service_type_name"`
	ExtraServiceName string    `json:"extra_service_name"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
}

type ReservationCompletedResponse struct {
	ReservationId    int       `json:"reservation_id"`
	ReservationDate  time.Time `json:"reservation_date"`
	DeviceTypeName   string    `json:"device_type_name"`
	BrandName        string    `json:"brand_name"`
	ModelName        string    `json:"model_name"`
	FixTypeName      string    `json:"fix_type_name"`
	ServiceTypeName  string    `json:"service_type_name"`
	ExtraServiceName string    `json:"extra_service_name"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
}

type ReservationPendingAndCompletedCountResponse struct {
	PendingCount   int64 `json:"pending_count"`
	CompletedCount int64 `json:"completed_count"`
}
