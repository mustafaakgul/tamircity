package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"time"
)

var expertiseService1ShiftMonday = &db.ExpertiseServiceShift{
	Day:          time.Monday,
	StartOfShift: 9,
	EndOfShift:   18,
	Status:       true,
}

var expertiseService1ShiftTuesday = &db.ExpertiseServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 10,
	EndOfShift:   17,
	Status:       true,
}

var expertiseService1ShiftWednesday = &db.ExpertiseServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   19,
	Status:       true,
}

var expertiseService1ShiftThursday = &db.ExpertiseServiceShift{
	Day:          time.Thursday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var expertiseService1ShiftFriday = &db.ExpertiseServiceShift{
	Day:          time.Friday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}

var expertiseService2ShiftMonday = &db.ExpertiseServiceShift{
	Day:          time.Monday,
	StartOfShift: 10,
	EndOfShift:   16,
	Status:       true,
}

var expertiseService2ShiftTuesday = &db.ExpertiseServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 8,
	EndOfShift:   19,
	Status:       true,
}

var expertiseService2ShiftWednesday = &db.ExpertiseServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 9,
	EndOfShift:   20,
	Status:       true,
}

var expertiseService2ShiftThursday = &db.ExpertiseServiceShift{
	Day:          time.Thursday,
	StartOfShift: 9,
	EndOfShift:   21,
	Status:       true,
}

var expertiseService2ShiftFriday = &db.ExpertiseServiceShift{
	Day:          time.Friday,
	StartOfShift: 12,
	EndOfShift:   22,
	Status:       true,
}

var expertiseService3ShiftMonday = &db.ExpertiseServiceShift{
	Day:          time.Monday,
	StartOfShift: 12,
	EndOfShift:   13,
	Status:       true,
}

var expertiseService3ShiftTuesday = &db.ExpertiseServiceShift{
	Day:          time.Tuesday,
	StartOfShift: 14,
	EndOfShift:   15,
	Status:       true,
}

var expertiseService3ShiftWednesday = &db.ExpertiseServiceShift{
	Day:          time.Wednesday,
	StartOfShift: 15,
	EndOfShift:   22,
	Status:       true,
}

var expertiseService3ShiftThursday = &db.ExpertiseServiceShift{
	Day:          time.Thursday,
	StartOfShift: 6,
	EndOfShift:   15,
	Status:       true,
}

var expertiseService3ShiftFriday = &db.ExpertiseServiceShift{
	Day:          time.Friday,
	StartOfShift: 10,
	EndOfShift:   19,
	Status:       true,
}
