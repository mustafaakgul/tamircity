package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"time"
)

var technicalService1ShiftMonday = &db.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 9,
	EndOfShift:   18,
	Status:       true,
}

var technicalService1ShiftTuesday = &db.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 10,
	EndOfShift:   17,
	Status:       true,
}

var technicalService1ShiftWednesday = &db.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   19,
	Status:       true,
}

var technicalService1ShiftThursday = &db.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var technicalService1ShiftFriday = &db.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var technicalService2ShiftMonday = &db.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 10,
	EndOfShift:   16,
	Status:       true,
}

var technicalService2ShiftTuesday = &db.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 8,
	EndOfShift:   19,
	Status:       true,
}

var technicalService2ShiftWednesday = &db.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   20,
	Status:       true,
}

var technicalService2ShiftThursday = &db.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 9,
	EndOfShift:   21,
	Status:       true,
}

var technicalService2ShiftFriday = &db.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 12,
	EndOfShift:   22,
	Status:       true,
}

var technicalService3ShiftMonday = &db.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 12,
	EndOfShift:   13,
	Status:       true,
}

var technicalService3ShiftTuesday = &db.TechnicalServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 14,
	EndOfShift:   15,
	Status:       true,
}

var technicalService3ShiftWednesday = &db.TechnicalServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 15,
	EndOfShift:   22,
	Status:       true,
}

var technicalService3ShiftThursday = &db.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 6,
	EndOfShift:   15,
	Status:       true,
}

var technicalService3ShiftFriday = &db.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 10,
	EndOfShift:   19,
	Status:       true,
}
