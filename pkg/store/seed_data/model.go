package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

// PC
var ModelSamsungPC = &tech_service.Model{
	Name:              "Samsung",
	ShortDescription:  "Samsung",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandSamsung,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelApplePC = &tech_service.Model{
	Name:              "Apple",
	ShortDescription:  "Apple",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelLenovoPC = &tech_service.Model{
	Name:              "Lenovo",
	ShortDescription:  "Lenovo",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandLenovo,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelHpPC = &tech_service.Model{
	Name:              "HP",
	ShortDescription:  "HP",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandHp,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelAsusPC = &tech_service.Model{
	Name:              "Asus",
	ShortDescription:  "Asus",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAsus,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelAcerPC = &tech_service.Model{
	Name:              "Acer",
	ShortDescription:  "Acer",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAcer,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelToshibaPC = &tech_service.Model{
	Name:              "Toshiba",
	ShortDescription:  "Toshiba",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandToshiba,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelDellPC = &tech_service.Model{
	Name:              "Dell",
	ShortDescription:  "Dell",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandDell,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelCasperPC = &tech_service.Model{
	Name:              "Casper",
	ShortDescription:  "Casper",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandCasper,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelHuaweiPC = &tech_service.Model{
	Name:              "Huawei",
	ShortDescription:  "Huawei",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandHuawei,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelMSIPC = &tech_service.Model{
	Name:              "MSI",
	ShortDescription:  "MSI",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMsi,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelMicrosoftPC = &tech_service.Model{
	Name:              "Microsoft",
	ShortDescription:  "Microsoft",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMicrosoft,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelMonsterPC = &tech_service.Model{
	Name:              "Monster",
	ShortDescription:  "Monster",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMonster,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelAsusRogPC = &tech_service.Model{
	Name:              "Asus Rog",
	ShortDescription:  "Asus Rog",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAsusRog,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

// Tablet LG
var ModelGpad7 = &tech_service.Model{
	Name:              "G Pad 7.0 V400",
	ShortDescription:  "G Pad 7.0 V400",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelGpad8 = &tech_service.Model{
	Name:              "G Pad 8.3 V500",
	ShortDescription:  "G Pad 8.3 V500",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelLGOther = &tech_service.Model{
	Name:              "Diger",
	ShortDescription:  "Diger",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

// Tablet General Mobile
var ModelGMTAB = &tech_service.Model{
	Name:              "e-Tab",
	ShortDescription:  "e-Tab",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelGMTAB4 = &tech_service.Model{
	Name:              "e-Tab 4",
	ShortDescription:  "e-Tab 4",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelGMTAB5 = &tech_service.Model{
	Name:              "e-Tab 5",
	ShortDescription:  "e-Tab 5",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelGMTAB10 = &tech_service.Model{
	Name:              "e-Tab 10",
	ShortDescription:  "e-Tab 10",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModelGMTAB8 = &tech_service.Model{
	Name:              "Discovery Tab8 3G",
	ShortDescription:  "Discovery Tab8 3G",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

///////////
var ModelIphone11 = &tech_service.Model{
	Name:              "iPhone 11",
	ShortDescription:  "iPhone 11",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelIphone12 = &tech_service.Model{
	Name:              "iPhone 12",
	ShortDescription:  "iPhone 12",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelIphone12Pro = &tech_service.Model{
	Name:              "iPhone 12 Pro",
	ShortDescription:  "iPhone 12 Pro",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1, TechnicalService2},
}

var ModelSamsungGalaxyS7 = &tech_service.Model{
	Name:              "Galaxy S7",
	ShortDescription:  "Galaxy S7",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandSamsung,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService3},
}

var ModelSamsungGalaxyS9 = &tech_service.Model{
	Name:              "Galaxy S9",
	ShortDescription:  "Galaxy S9",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandSamsung,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService3},
}

var ModelLenovoM7 = &tech_service.Model{
	Name:              "Lenovo Tab M7",
	ShortDescription:  "Lenovo Tab M7",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLenovo,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService2, TechnicalService3},
}

var ModelLenovoM8 = &tech_service.Model{
	Name:              "Lenovo Tab M8",
	ShortDescription:  "Lenovo Tab M8",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLenovo,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService2, TechnicalService3},
}

var ModeliPad6 = &tech_service.Model{
	Name:              "iPad 6.Nesil",
	ShortDescription:  "iPad 6.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}

var ModeliPad9 = &tech_service.Model{
	Name:              "iPad 9.Nesil",
	ShortDescription:  "iPad 9.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandApple,
	TechnicalServices: []*tech_service.TechnicalService{TechnicalService1},
}
