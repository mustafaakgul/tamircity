package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var ModelIphone11 = &db.Model{
	Name:              "iPhone 11",
	ShortDescription:  "iPhone 11",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandsApple,
	TechnicalServices: []*db.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelIphone12 = &db.Model{
	Name:              "iPhone 12",
	ShortDescription:  "iPhone 12",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandsApple,
	TechnicalServices: []*db.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelIphone12Pro = &db.Model{
	Name:              "iPhone 12 Pro",
	ShortDescription:  "iPhone 12 Pro",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandsApple,
	TechnicalServices: []*db.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelSamsungGalaxyS7 = &db.Model{
	Name:              "Galaxy S7",
	ShortDescription:  "Galaxy S7",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandsSamsung,
	TechnicalServices: []*db.TechnicalService{TechnicalService3},
}

var ModelSamsungGalaxyS9 = &db.Model{
	Name:              "Galaxy S9",
	ShortDescription:  "Galaxy S9",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandsSamsung,
	TechnicalServices: []*db.TechnicalService{TechnicalService3},
}

var ModelLenovoM7 = &db.Model{
	Name:              "Lenovo Tab M7",
	ShortDescription:  "Lenovo Tab M7",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandsLenovo,
	TechnicalServices: []*db.TechnicalService{TechnicalService2, TechnicalService3},
}

var ModelLenovoM8 = &db.Model{
	Name:              "Lenovo Tab M8",
	ShortDescription:  "Lenovo Tab M8",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandsLenovo,
	TechnicalServices: []*db.TechnicalService{TechnicalService2, TechnicalService3},
}

var ModeliPad6 = &db.Model{
	Name:              "iPad 6.Nesil",
	ShortDescription:  "iPad 6.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandsApple,
	TechnicalServices: []*db.TechnicalService{TechnicalService1},
}

var ModeliPad9 = &db.Model{
	Name:              "iPad 9.Nesil",
	ShortDescription:  "iPad 9.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandsApple,
	TechnicalServices: []*db.TechnicalService{TechnicalService1},
}
