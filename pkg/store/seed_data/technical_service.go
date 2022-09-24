package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var TechnicalService1 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 1",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Basaksehir/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalService1ShiftMonday, technicalService1ShiftTuesday, technicalService1ShiftWednesday, technicalService1ShiftThursday, technicalService1ShiftFriday},
}

var TechnicalService2 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 2",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Bahcelievler/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalService2ShiftMonday, technicalService2ShiftTuesday, technicalService2ShiftWednesday, technicalService2ShiftThursday, technicalService2ShiftFriday},
}

var TechnicalService3 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 3",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Address:                "Beylikduzu/ISTANBUL",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{DeviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalService3ShiftMonday, technicalService3ShiftTuesday, technicalService3ShiftWednesday, technicalService3ShiftThursday, technicalService3ShiftFriday},
}
