package postgres

import (
	dbModels "github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {

	// TODO: Develop a better way to do HasTable
	if !db.Migrator().HasTable(&tech_service.Brand{}) {
		log.Printf("Migrations started")
		db.AutoMigrate(
			&tech_service.Brand{},
			&dbModels.Contact{},
			&dbModels.Customer{},
			&tech_service.DeviceType{},
			&tech_service.ExtraService{},
			&tech_service.FixType{},
			&tech_service.Model{},
			&dbModels.Newsletter{},
			&tech_service.Reservation{},
			&tech_service.ServiceType{},
			&tech_service.TechnicalService{},
			&tech_service.TechnicalServiceCandidate{},
			&dbModels.User{},
			&tech_service.TechnicalServiceReservation{},
			&tech_service.TechnicalServiceShift{},
		)

		Seeder(db)

		log.Println("Migrations done")
	}
	return error(nil)
}
