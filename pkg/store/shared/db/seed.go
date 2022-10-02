package postgres

import (
	"github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	// Adding Seed data
	technicalServiceStore := tech_service.NewTechnicalServiceStore(db)
	brandStore := tech_service.NewBrandStore(db)
	modelStore := tech_service.NewModelStore(db)
	fixTypeStore := tech_service.NewFixTypeStore(db)
	serviceTypeStore := tech_service.NewServiceTypeStore(db)
	extraServiceStore := tech_service.NewExtraServiceStore(db)
	reservationStore := tech_service.NewReservationStore(db)
	/*deviceTypeStore := repositories.NewDeviceTypeStore(db)
	 */
	technicalServiceStore.Seed()
	brandStore.Seed()
	modelStore.Seed()
	fixTypeStore.Seed()
	serviceTypeStore.Seed()
	extraServiceStore.Seed()
	reservationStore.Seed()

	log.Printf("Seeding done")

	return error(nil)
}
