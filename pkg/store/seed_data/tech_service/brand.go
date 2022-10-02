package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

var BrandApple = &tech_service.Brand{
	Name:        "Apple",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var BrandSamsung = &tech_service.Brand{
	Name:        "Samsung",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var BrandLenovo = &tech_service.Brand{
	Name:        "Lenovo",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var BrandHp = &tech_service.Brand{
	Name:        "HP",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandAsus = &tech_service.Brand{
	Name:        "Asus",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone},
}

var BrandAcer = &tech_service.Brand{
	Name:        "Acer",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandToshiba = &tech_service.Brand{
	Name:        "Toshiba",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandDell = &tech_service.Brand{
	Name:        "Dell",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandCasper = &tech_service.Brand{
	Name:        "Casper",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypeTablet, DeviceTypePhone},
}

var BrandHuawei = &tech_service.Brand{
	Name:        "Huawei",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypeTablet, DeviceTypePhone},
}

var BrandMsi = &tech_service.Brand{
	Name:        "MSI",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandMicrosoft = &tech_service.Brand{
	Name:        "Microsoft",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypeTablet},
}

var BrandMonster = &tech_service.Brand{
	Name:        "Monster",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandAsusRog = &tech_service.Brand{
	Name:        "Asus Rog",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc},
}

var BrandNokia = &tech_service.Brand{
	Name:        "Nokia",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandOppo = &tech_service.Brand{
	Name:        "Oppo",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}
var BrandGeneralMobile = &tech_service.Brand{
	Name:        "General Mobile",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone, DeviceTypeTablet},
}

var BrandHometech = &tech_service.Brand{
	Name:        "Hometech",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet},
}

var BrandXiaomi = &tech_service.Brand{
	Name:        "Xiaomi",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet, DeviceTypePhone},
}

var BrandAlcatel = &tech_service.Brand{
	Name:        "Alcatel",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet, DeviceTypePhone},
}

var BrandReeder = &tech_service.Brand{
	Name:        "Reeder",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet, DeviceTypePhone},
}

var BrandPhilips = &tech_service.Brand{
	Name:        "Philips",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet},
}

var BrandAmazon = &tech_service.Brand{
	Name:        "Amazon",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet},
}

var BrandLG = &tech_service.Brand{
	Name:        "LG",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypeTablet, DeviceTypePhone},
}

var BrandBlackberry = &tech_service.Brand{
	Name:        "Blackberry",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandGoogle = &tech_service.Brand{
	Name:        "Google",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandHiking = &tech_service.Brand{
	Name:        "Hiking",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandHonor = &tech_service.Brand{
	Name:        "Honor",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandHTC = &tech_service.Brand{
	Name:        "HTC",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandInfinix = &tech_service.Brand{
	Name:        "Infinix",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandOnePlus = &tech_service.Brand{
	Name:        "OnePlus",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandRealme = &tech_service.Brand{
	Name:        "Realme",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandSony = &tech_service.Brand{
	Name:        "Sony",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandTCL = &tech_service.Brand{
	Name:        "TCL",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandTecno = &tech_service.Brand{
	Name:        "Tecno",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandVestel = &tech_service.Brand{
	Name:        "Vestel",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}

var BrandVivo = &tech_service.Brand{
	Name:        "Vivo",
	IsActive:    true,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePhone},
}
