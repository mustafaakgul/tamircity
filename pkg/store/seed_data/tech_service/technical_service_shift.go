package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"time"
)

var technicalService1ShiftMonday = &tech_service.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 9,
	EndOfShift:   18,
	Status:       true,
}

var technicalService1ShiftTuesday = &tech_service.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 10,
	EndOfShift:   17,
	Status:       true,
}

var technicalService1ShiftWednesday = &tech_service.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   19,
	Status:       true,
}

var technicalService1ShiftThursday = &tech_service.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var technicalService1ShiftFriday = &tech_service.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var technicalService2ShiftMonday = &tech_service.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 10,
	EndOfShift:   16,
	Status:       true,
}

var technicalService2ShiftTuesday = &tech_service.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 8,
	EndOfShift:   19,
	Status:       true,
}

var technicalService2ShiftWednesday = &tech_service.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   20,
	Status:       true,
}

var technicalService2ShiftThursday = &tech_service.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 9,
	EndOfShift:   21,
	Status:       true,
}

var technicalService2ShiftFriday = &tech_service.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 12,
	EndOfShift:   22,
	Status:       true,
}

var technicalService3ShiftMonday = &tech_service.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 12,
	EndOfShift:   13,
	Status:       true,
}

var technicalService3ShiftTuesday = &tech_service.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 14,
	EndOfShift:   15,
	Status:       true,
}

var technicalService3ShiftWednesday = &tech_service.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 15,
	EndOfShift:   22,
	Status:       true,
}

var technicalService3ShiftThursday = &tech_service.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 6,
	EndOfShift:   15,
	Status:       true,
}

var technicalService3ShiftFriday = &tech_service.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 10,
	EndOfShift:   19,
	Status:       true,
}
