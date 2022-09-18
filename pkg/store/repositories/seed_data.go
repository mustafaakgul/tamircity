package repositories

import "github.com/mustafakocatepe/Tamircity/pkg/models/db"

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

var modelIphone11 = &db.Model{
	Name:             "iPhone 11",
	ShortDescription: "iPhone 11",
	IsActive:         true,
	DeviceType:       deviceTypePhone,
	Brand:            brandsApple,
}

var modelIphone12 = &db.Model{
	Name:             "iPhone 12",
	ShortDescription: "iPhone 12",
	IsActive:         true,
	DeviceType:       deviceTypePhone,
	Brand:            brandsApple,
}

var modelIphone12Pro = &db.Model{
	Name:             "iPhone 12 Pro",
	ShortDescription: "iPhone 12 Pro",
	IsActive:         true,
	DeviceType:       deviceTypePhone,
	Brand:            brandsApple,
}

var modelSamsungGalaxyS7 = &db.Model{
	Name:             "Galaxy S7",
	ShortDescription: "Galaxy S7",
	IsActive:         true,
	DeviceType:       deviceTypePhone,
	Brand:            brandsSamsung,
}

var modelSamsungGalaxyS9 = &db.Model{
	Name:             "Galaxy S9",
	ShortDescription: "Galaxy S9",
	IsActive:         true,
	DeviceType:       deviceTypePhone,
	Brand:            brandsSamsung,
}

var modelLenovoM7 = &db.Model{
	Name:             "Lenovo Tab M7",
	ShortDescription: "Lenovo Tab M7",
	IsActive:         true,
	DeviceType:       deviceTypeTablet,
	Brand:            brandsLenovo,
}

var modelLenovoM8 = &db.Model{
	Name:             "Lenovo Tab M8",
	ShortDescription: "Lenovo Tab M8",
	IsActive:         true,
	DeviceType:       deviceTypeTablet,
	Brand:            brandsLenovo,
}

var modeliPad6 = &db.Model{
	Name:             "iPad 6.Nesil",
	ShortDescription: "iPad 6.Nesil",
	IsActive:         true,
	DeviceType:       deviceTypeTablet,
	Brand:            brandsApple,
}

var modeliPad9 = &db.Model{
	Name:             "iPad 9.Nesil",
	ShortDescription: "iPad 9.Nesil",
	IsActive:         true,
	DeviceType:       deviceTypeTablet,
	Brand:            brandsApple,
}
