package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
)

// PC
var ModelSamsungPC = &db.Model{
	Name:              "Samsung",
	ShortDescription:  "Samsung",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandSamsung,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var ModelApplePC = &db.Model{
	Name:              "Apple",
	ShortDescription:  "Apple",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService2},
}

var ModelLenovoPC = &db.Model{
	Name:              "Lenovo",
	ShortDescription:  "Lenovo",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandLenovo,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService2},
}

var ModelHpPC = &db.Model{
	Name:              "HP",
	ShortDescription:  "HP",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandHp,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelAsusPC = &db.Model{
	Name:              "Asus",
	ShortDescription:  "Asus",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAsus,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelAcerPC = &db.Model{
	Name:              "Acer",
	ShortDescription:  "Acer",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAcer,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelToshibaPC = &db.Model{
	Name:              "Toshiba",
	ShortDescription:  "Toshiba",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandToshiba,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelDellPC = &db.Model{
	Name:              "Dell",
	ShortDescription:  "Dell",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandDell,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelCasperPC = &db.Model{
	Name:              "Casper",
	ShortDescription:  "Casper",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandCasper,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelHuaweiPC = &db.Model{
	Name:              "Huawei",
	ShortDescription:  "Huawei",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandHuawei,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelMSIPC = &db.Model{
	Name:              "MSI",
	ShortDescription:  "MSI",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMsi,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelMicrosoftPC = &db.Model{
	Name:              "Microsoft",
	ShortDescription:  "Microsoft",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMicrosoft,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelMonsterPC = &db.Model{
	Name:              "Monster",
	ShortDescription:  "Monster",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandMonster,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelAsusRogPC = &db.Model{
	Name:              "Asus Rog",
	ShortDescription:  "Asus Rog",
	IsActive:          true,
	DeviceType:        DeviceTypePc,
	Brand:             BrandAsusRog,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

// Tablet LG
var ModelGpad7 = &db.Model{
	Name:              "G Pad 7.0 V400",
	ShortDescription:  "G Pad 7.0 V400",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelGpad8 = &db.Model{
	Name:              "G Pad 8.3 V500",
	ShortDescription:  "G Pad 8.3 V500",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelLGOther = &db.Model{
	Name:              "Diger",
	ShortDescription:  "Diger",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLG,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

// Tablet General Mobile
var ModelGMTAB = &db.Model{
	Name:              "e-Tab",
	ShortDescription:  "e-Tab",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelGMTAB4 = &db.Model{
	Name:              "e-Tab 4",
	ShortDescription:  "e-Tab 4",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelGMTAB5 = &db.Model{
	Name:              "e-Tab 5",
	ShortDescription:  "e-Tab 5",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelGMTAB10 = &db.Model{
	Name:              "e-Tab 10",
	ShortDescription:  "e-Tab 10",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelGMTAB8 = &db.Model{
	Name:              "Discovery Tab8 3G",
	ShortDescription:  "Discovery Tab8 3G",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandGeneralMobile,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

///////////
var ModelIphone11 = &db.Model{
	Name:              "iPhone 11",
	ShortDescription:  "iPhone 11",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelIphone12 = &db.Model{
	Name:              "iPhone 12",
	ShortDescription:  "iPhone 12",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelIphone12Pro = &db.Model{
	Name:              "iPhone 12 Pro",
	ShortDescription:  "iPhone 12 Pro",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelSamsungGalaxyS7 = &db.Model{
	Name:              "Galaxy S7",
	ShortDescription:  "Galaxy S7",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandSamsung,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelSamsungGalaxyS9 = &db.Model{
	Name:              "Galaxy S9",
	ShortDescription:  "Galaxy S9",
	IsActive:          true,
	DeviceType:        DeviceTypePhone,
	Brand:             BrandSamsung,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelLenovoM7 = &db.Model{
	Name:              "Lenovo Tab M7",
	ShortDescription:  "Lenovo Tab M7",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLenovo,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModelLenovoM8 = &db.Model{
	Name:              "Lenovo Tab M8",
	ShortDescription:  "Lenovo Tab M8",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandLenovo,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModeliPad6 = &db.Model{
	Name:              "iPad 6.Nesil",
	ShortDescription:  "iPad 6.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}

var ModeliPad9 = &db.Model{
	Name:              "iPad 9.Nesil",
	ShortDescription:  "iPad 9.Nesil",
	IsActive:          true,
	DeviceType:        DeviceTypeTablet,
	Brand:             BrandApple,
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1},
}
