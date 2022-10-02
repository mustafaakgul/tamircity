package tech_service

import "time"

type TechnicalServiceShiftRequest struct {
	Day          time.Weekday `json:"day"`
	StartOfShift string       `json:"start_of_shift"`
	EndOfShift   string       `json:"end_of_shift"`
}

type TechnicalServiceShiftResponse struct {
	Id           uint         `json:"id"`
	Day          time.Weekday `json:"day"`
	StartOfShift string       `json:"start_of_shift"`
	EndOfShift   string       `json:"end_of_shift"`
}
