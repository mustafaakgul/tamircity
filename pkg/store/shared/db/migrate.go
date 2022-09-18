package postgres

import (
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {

	// TODO: Develop a better way to do HasTable
	if !db.Migrator().HasTable(&dbModels.Brand{}) {
		log.Printf("Migrations started")
		db.AutoMigrate(
			&dbModels.Brand{},
			&dbModels.Contact{},
			&dbModels.Customer{},
			&dbModels.DeviceType{},
			&dbModels.ExtraService{},
			&dbModels.FixType{},
			&dbModels.Model{},
			&dbModels.Newsletter{},
			&dbModels.Reservation{},
			&dbModels.ServiceType{},
			&dbModels.TechnicalService{},
			&dbModels.TechnicalServiceReservation{},
			&dbModels.TechnicalServiceShift{},
		)
		log.Println("Migrations done")
	}
	return error(nil)
}
