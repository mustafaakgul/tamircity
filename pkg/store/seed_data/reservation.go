package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"time"
)

var Reservation1 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       13,
	EndOfHour:         14,
	Price:             100,
	Status:            db.Pending,
	FullName:          "Kullanıcı 1",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelIphone12,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}

var Reservation2 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       15,
	EndOfHour:         16,
	Price:             150,
	Status:            db.Cancelled,
	FullName:          "Kullanıcı 2",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelIphone12Pro,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}

var Reservation3 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       17,
	EndOfHour:         18,
	Price:             150,
	Status:            db.Pending,
	FullName:          "Kullanıcı 3",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelIphone11,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}

var Reservation4 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       12,
	EndOfHour:         13,
	Price:             200,
	Status:            db.Completed,
	FullName:          "Kullanıcı 4",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelLenovoM7,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}

var Reservation5 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       15,
	EndOfHour:         16,
	Price:             400,
	Status:            db.Approved,
	FullName:          "Kullanıcı 5",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelLenovoM7,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}

var Reservation6 = &db.Reservation{
	Date:              time.Now(),
	StartOfHour:       9,
	EndOfHour:         10,
	Price:             1000,
	Status:            db.Approved,
	FullName:          "Kullanıcı 6",
	Email:             "msk@test.com",
	PhoneNumber:       "05555555555",
	SecondPhoneNumber: "05555555555",
	Description:       "Güzel hizmet",
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ModelEntity:       ModelLenovoM7,
	ServiceType:       ServiceType1,
	ExpertiseService:  ExpertiseService1,
}
