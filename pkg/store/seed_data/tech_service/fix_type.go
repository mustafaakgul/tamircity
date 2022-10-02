package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

var FixType1 = &tech_service.FixType{
	Description: "Ekran Değişimi",
	IsActive:    true,
	Price:       400,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType2 = &tech_service.FixType{
	Description: "Batarya Değişimi",
	IsActive:    true,
	Price:       1000,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType3 = &tech_service.FixType{
	Description: "Açma Kapama Tuşu Sorunu",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType4 = &tech_service.FixType{
	Description: "Hoparlör",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*tech_service.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}
