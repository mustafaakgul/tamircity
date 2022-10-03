package postgres

import (
	dbModels "github.com/anthophora/tamircity/pkg/models/db"
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
			&dbModels.Model{},
			&dbModels.Newsletter{},
			&dbModels.Reservation{},
			&dbModels.ServiceType{},
			&dbModels.ExpertiseService{},
			&dbModels.ExpertiseServiceCandidate{},
			&dbModels.User{},
			&dbModels.ExpertiseServiceShift{},
		)

		Seeder(db)
		log.Println("Migrations done")
	}
	return error(nil)
}
