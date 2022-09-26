package web

import "time"

type TechnicalServiceShiftRequest struct {
	Day          time.Weekday `json:"day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}

type TechnicalServiceShiftResponse struct {
	Id           uint         `json:"id"`
	Day          time.Weekday `json:"day"`
	StartOfShift int          `json:"start_of_shift"`
	EndOfShift   int          `json:"end_of_shift"`
}
