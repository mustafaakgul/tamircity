package web

import "time"

type ExpertiseServiceShiftRequest struct {
	Day          time.Weekday `json:"day"`
	StartOfShift string       `json:"start_of_shift"`
	EndOfShift   string       `json:"end_of_shift"`
}

type ExpertiseServiceShiftResponse struct {
	Id           uint         `json:"id"`
	Day          time.Weekday `json:"day"`
	StartOfShift string       `json:"start_of_shift"`
	EndOfShift   string       `json:"end_of_shift"`
}
