package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
)

var BrandApple = &db.Brand{
	Name:              "Apple",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandSamsung = &db.Brand{
	Name:              "Samsung",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandLenovo = &db.Brand{
	Name:              "Lenovo",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHp = &db.Brand{
	Name:              "HP",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandAsus = &db.Brand{
	Name:              "Asus",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandAcer = &db.Brand{
	Name:              "Acer",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandToshiba = &db.Brand{
	Name:              "Toshiba",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandDell = &db.Brand{
	Name:              "Dell",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandCasper = &db.Brand{
	Name:              "Casper",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHuawei = &db.Brand{
	Name:              "Huawei",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandMsi = &db.Brand{
	Name:              "MSI",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandMicrosoft = &db.Brand{
	Name:              "Microsoft",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc, DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandMonster = &db.Brand{
	Name:              "Monster",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandAsusRog = &db.Brand{
	Name:              "Asus Rog",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePc},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandNokia = &db.Brand{
	Name:              "Nokia",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandOppo = &db.Brand{
	Name:              "Oppo",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}
var BrandGeneralMobile = &db.Brand{
	Name:              "General Mobile",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone, DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHometech = &db.Brand{
	Name:              "Hometech",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandXiaomi = &db.Brand{
	Name:              "Xiaomi",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandAlcatel = &db.Brand{
	Name:              "Alcatel",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandReeder = &db.Brand{
	Name:              "Reeder",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandPhilips = &db.Brand{
	Name:              "Philips",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandAmazon = &db.Brand{
	Name:              "Amazon",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandLG = &db.Brand{
	Name:              "LG",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypeTablet, DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandBlackberry = &db.Brand{
	Name:              "Blackberry",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandGoogle = &db.Brand{
	Name:              "Google",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHiking = &db.Brand{
	Name:              "Hiking",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHonor = &db.Brand{
	Name:              "Honor",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandHTC = &db.Brand{
	Name:              "HTC",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandInfinix = &db.Brand{
	Name:              "Infinix",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandOnePlus = &db.Brand{
	Name:              "OnePlus",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandRealme = &db.Brand{
	Name:              "Realme",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandSony = &db.Brand{
	Name:              "Sony",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandTCL = &db.Brand{
	Name:              "TCL",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandTecno = &db.Brand{
	Name:              "Tecno",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandVestel = &db.Brand{
	Name:              "Vestel",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}

var BrandVivo = &db.Brand{
	Name:              "Vivo",
	IsActive:          true,
	DeviceTypes:       []*db.DeviceType{DeviceTypePhone},
	ExpertiseServices: []*db.ExpertiseService{ExpertiseService1, ExpertiseService2},
}
