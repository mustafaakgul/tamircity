package tech_service

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
)

var ExtraServiceFirst = &tech_service.ExtraService{
	Description: "Cihaz Bakım Paketi",
	Price:       150,
}

var ExtraServiceSecond = &tech_service.ExtraService{
	Description: "Cihaz Temizlik Paketi",
	Price:       75,
}

var ExtraServiceThird = &tech_service.ExtraService{
	Description: "Cihaz Kılıfı",
	Price:       50,
}

var ExtraServiceFourth = &tech_service.ExtraService{
	Description: "Cihaz Ekran Koruma",
	Price:       25,
}
