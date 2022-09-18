package repositories

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"time"
)

var deviceTypePc = &db.DeviceType{
	Name:             "Personel Computer",
	ShortDescription: "PC",
	IsActive:         true,
}

var deviceTypePhone = &db.DeviceType{
	Name:             "Phone",
	ShortDescription: "Phone",
	IsActive:         true,
}

var deviceTypeTablet = &db.DeviceType{
	Name:             "Tablet",
	ShortDescription: "Tablet",
	IsActive:         true,
}

var brandsApple = &db.Brand{
	Name:        "Apple",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

var brandsSamsung = &db.Brand{
	Name:        "Samsung",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

var brandsLenovo = &db.Brand{
	Name:        "Lenovo",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePhone, deviceTypeTablet},
}

var brandNokia = &db.Brand{
	Name:        "Nokia",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePhone},
}
var brandOppo = &db.Brand{
	Name:        "Oppo",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePhone},
}
var brandGeneralMobile = &db.Brand{
	Name:        "General Mobile",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypePhone},
}

var brandHometech = &db.Brand{
	Name:        "Hometech",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{deviceTypeTablet},
}

var fixType1 = &db.FixType{
	Description: "Ekran Değişimi",
	IsActive:    true,
	Price:       400,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

var fixType2 = &db.FixType{
	Description: "Batarya Değişimi",
	IsActive:    true,
	Price:       1000,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

var fixType3 = &db.FixType{
	Description: "Açma Kapama Tuşu Sorunu",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

var fixType4 = &db.FixType{
	Description: "Hoparlör",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{deviceTypePc, deviceTypePhone, deviceTypeTablet},
}

//Models
var modelIphone11 = &db.Model{
	Name:              "iPhone 11",
	ShortDescription:  "iPhone 11",
	IsActive:          true,
	DeviceType:        deviceTypePhone,
	Brand:             brandsApple,
	TechnicalServices: []*db.TechnicalService{technicalService1, technicalService2},
}

var modelIphone12 = &db.Model{
	Name:              "iPhone 12",
	ShortDescription:  "iPhone 12",
	IsActive:          true,
	DeviceType:        deviceTypePhone,
	Brand:             brandsApple,
	TechnicalServices: []*db.TechnicalService{technicalService1, technicalService2},
}

var modelIphone12Pro = &db.Model{
	Name:              "iPhone 12 Pro",
	ShortDescription:  "iPhone 12 Pro",
	IsActive:          true,
	DeviceType:        deviceTypePhone,
	Brand:             brandsApple,
	TechnicalServices: []*db.TechnicalService{technicalService1, technicalService2},
}

var modelSamsungGalaxyS7 = &db.Model{
	Name:              "Galaxy S7",
	ShortDescription:  "Galaxy S7",
	IsActive:          true,
	DeviceType:        deviceTypePhone,
	Brand:             brandsSamsung,
	TechnicalServices: []*db.TechnicalService{technicalService3},
}

var modelSamsungGalaxyS9 = &db.Model{
	Name:              "Galaxy S9",
	ShortDescription:  "Galaxy S9",
	IsActive:          true,
	DeviceType:        deviceTypePhone,
	Brand:             brandsSamsung,
	TechnicalServices: []*db.TechnicalService{technicalService3},
}

var modelLenovoM7 = &db.Model{
	Name:              "Lenovo Tab M7",
	ShortDescription:  "Lenovo Tab M7",
	IsActive:          true,
	DeviceType:        deviceTypeTablet,
	Brand:             brandsLenovo,
	TechnicalServices: []*db.TechnicalService{technicalService2, technicalService3},
}

var modelLenovoM8 = &db.Model{
	Name:              "Lenovo Tab M8",
	ShortDescription:  "Lenovo Tab M8",
	IsActive:          true,
	DeviceType:        deviceTypeTablet,
	Brand:             brandsLenovo,
	TechnicalServices: []*db.TechnicalService{technicalService2, technicalService3},
}

var modeliPad6 = &db.Model{
	Name:              "iPad 6.Nesil",
	ShortDescription:  "iPad 6.Nesil",
	IsActive:          true,
	DeviceType:        deviceTypeTablet,
	Brand:             brandsApple,
	TechnicalServices: []*db.TechnicalService{technicalService1},
}

var modeliPad9 = &db.Model{
	Name:              "iPad 9.Nesil",
	ShortDescription:  "iPad 9.Nesil",
	IsActive:          true,
	DeviceType:        deviceTypeTablet,
	Brand:             brandsApple,
	TechnicalServices: []*db.TechnicalService{technicalService1},
}

var serviceType1 = &db.ServiceType{
	Name:        "Yerinde Tamir",
	Description: "İstanbul içinde bulunduğunuz yere geliyoruz. Mobil servis aracımızda cihazınızın tamirini gerçekleştiriyoruz. Covid-19 prosedürü kapsamında kapıdan teslim alıp, garanti belgesi ile teslim ediyoruz.",
	Price:       200,
}

var serviceType2 = &db.ServiceType{
	Name:        "Merkezde Tamir",
	Description: "Servisimize gelmek isteyen kullanıcılarımız için, kampüsümüzde ister bahçede ister ferah kapalı bekleme salonumuzda vakit geçirme imkânı bulunmaktadır.",
	Price:       0,
}
var serviceType3 = &db.ServiceType{
	Name:        "Kargo",
	Description: "Dilediğiniz kargo firması ile merkez servisimize arızalı telefonunuzu gönderebilirsiniz. Teslim etmeden önce, sim kartınızı ve kılıfınızı alarak, varsa kendi orijinal kutusunda, yoksa korunaklı şekilde paketleyiniz",
	Price:       0,
}

var serviceType4 = &db.ServiceType{
	Name:        "Kurye",
	Description: "İstanbul'da kuryemizle bulunduğunuz yerden cihazınızı teslim alıyor, tamir sonrası dilediğiniz yere kurye ile teslim ediyoruz. Teslim etmeden önce, sim kartınızı ve kılıfınızı alarak, varsa kendi orijinal kutusunda, yoksa korunaklı şekilde paketleyiniz.",
	Price:       75,
}

//Technical Service
var technicalService1 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 1",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{deviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalServiceShift1, technicalServiceShift2, technicalServiceShift3},
}

var technicalService2 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 2",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{deviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalServiceShift4, technicalServiceShift2, technicalServiceShift3},
}

var technicalService3 = &db.TechnicalService{

	ServiceName:            "Teknik Servis 3",
	IdentityNumber:         "123456789",
	PhoneNumber:            "123456789",
	Email:                  "email1@email.com",
	Iban:                   "123456789",
	IsActive:               true,
	DeviceTypes:            []*db.DeviceType{deviceTypePc},
	TechnicalServiceShifts: []*db.TechnicalServiceShift{technicalServiceShift4, technicalServiceShift2, technicalServiceShift3},
}

//Technical Service Shifts
var technicalServiceShift1 = &db.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 9,
	EndOfShift:   18,
	Status:       true,
}

var technicalServiceShift4 = &db.TechnicalServiceShift{
	Day:          time.Monday,
	StartOfShift: 10,
	EndOfShift:   17,
	Status:       true,
}

var technicalServiceShift2 = &db.TechnicalServiceShift{
	Day:          time.Friday,
	StartOfShift: 9,
	EndOfShift:   19,
	Status:       true,
}

var technicalServiceShift3 = &db.TechnicalServiceShift{
	Day:          time.Thursday,
	StartOfShift: 8,
	EndOfShift:   16,
	Status:       true,
}
