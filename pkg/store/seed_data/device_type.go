package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

var DeviceTypePc = &tech_service.DeviceType{
	Name:             "Personel Computer",
	ShortDescription: "PC",
	IsActive:         true,
}

var DeviceTypePhone = &tech_service.DeviceType{
	Name:             "Phone",
	ShortDescription: "Phone",
	IsActive:         true,
}

var DeviceTypeTablet = &tech_service.DeviceType{
	Name:             "Tablet",
	ShortDescription: "Tablet",
	IsActive:         true,
}
