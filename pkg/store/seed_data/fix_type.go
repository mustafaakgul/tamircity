package seed_data

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

var FixType1 = &tech_service.FixType{
	Description: "Ekran Değişimi",
	IsActive:    true,
	Price:       400,
	DeviceTypes: []*db.DeviceType{},
}

var FixType2 = &tech_service.FixType{
	Description: "Batarya Değişimi",
	IsActive:    true,
	Price:       1000,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType3 = &tech_service.FixType{
	Description: "Açma Kapama Tuşu Sorunu",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType4 = &tech_service.FixType{
	Description: "Hoparlör",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}
