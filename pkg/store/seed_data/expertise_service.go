package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
)

var ExpertiseService1 = &db.ExpertiseService{

	ServiceName:            "Teknik Servis 1",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Basaksehir/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	ExpertiseServiceShifts: []*db.ExpertiseServiceShift{expertiseService1ShiftMonday, expertiseService1ShiftTuesday, expertiseService1ShiftWednesday, expertiseService1ShiftThursday, expertiseService1ShiftFriday},
}

var ExpertiseService2 = &db.ExpertiseService{

	ServiceName:            "Teknik Servis 2",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Bahcelievler/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	ExpertiseServiceShifts: []*db.ExpertiseServiceShift{expertiseService2ShiftMonday, expertiseService2ShiftTuesday, expertiseService2ShiftWednesday, expertiseService2ShiftThursday, expertiseService2ShiftFriday},
}

var ExpertiseService3 = &db.ExpertiseService{

	ServiceName:            "Teknik Servis 3",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Beylikduzu/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	ExpertiseServiceShifts: []*db.ExpertiseServiceShift{expertiseService3ShiftMonday, expertiseService3ShiftTuesday, expertiseService3ShiftWednesday, expertiseService3ShiftThursday, expertiseService3ShiftFriday},
}
