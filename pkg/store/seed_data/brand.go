package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var BrandsApple = &db.Brand{
	Name:        "Apple",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var BrandsSamsung = &db.Brand{
	Name:        "Samsung",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var BrandsLenovo = &db.Brand{
	Name:        "Lenovo",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePhone, DeviceTypeTablet},
}

var BrandNokia = &db.Brand{
	Name:        "Nokia",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePhone},
}
var BrandOppo = &db.Brand{
	Name:        "Oppo",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePhone},
}
var BrandGeneralMobile = &db.Brand{
	Name:        "General Mobile",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypePhone},
}

var BrandHometech = &db.Brand{
	Name:        "Hometech",
	IsActive:    true,
	DeviceTypes: []*db.DeviceType{DeviceTypeTablet},
}
