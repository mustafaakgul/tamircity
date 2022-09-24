package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var FixType1 = &db.FixType{
	Description: "Ekran Değişimi",
	IsActive:    true,
	Price:       400,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType2 = &db.FixType{
	Description: "Batarya Değişimi",
	IsActive:    true,
	Price:       1000,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType3 = &db.FixType{
	Description: "Açma Kapama Tuşu Sorunu",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}

var FixType4 = &db.FixType{
	Description: "Hoparlör",
	IsActive:    true,
	Price:       250,
	DeviceTypes: []*db.DeviceType{DeviceTypePc, DeviceTypePhone, DeviceTypeTablet},
}
