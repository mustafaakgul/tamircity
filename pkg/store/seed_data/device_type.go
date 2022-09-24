package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var DeviceTypePc = &db.DeviceType{
	Name:             "Personel Computer",
	ShortDescription: "PC",
	IsActive:         true,
}

var DeviceTypePhone = &db.DeviceType{
	Name:             "Phone",
	ShortDescription: "Phone",
	IsActive:         true,
}

var DeviceTypeTablet = &db.DeviceType{
	Name:             "Tablet",
	ShortDescription: "Tablet",
	IsActive:         true,
}
