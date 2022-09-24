package seed_data

import "github.com/anthophora/tamircity/pkg/models/db"

var ExtraServiceFirst = &db.ExtraService{
	Description: "Cihaz Bakım Paketi",
	Price:       150,
}

var ExtraServiceSecond = &db.ExtraService{
	Description: "Cihaz Temizlik Paketi",
	Price:       75,
}

var ExtraServiceThird = &db.ExtraService{
	Description: "Cihaz Kılıfı",
	Price:       50,
}

var ExtraServiceFourth = &db.ExtraService{
	Description: "Cihaz Ekran Koruma",
	Price:       25,
}
