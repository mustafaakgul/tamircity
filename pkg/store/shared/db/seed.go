package postgres

import (
	"github.com/anthophora/tamircity/pkg/store/repositories"
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	// Adding Seed data
	technicalServiceStore := repositories.NewExpertiseServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	reservationStore := repositories.NewReservationStore(db)
	/*deviceTypeStore := repositories.NewDeviceTypeStore(db)
	 */
	technicalServiceStore.Seed()
	brandStore.Seed()
	modelStore.Seed()
	serviceTypeStore.Seed()
	reservationStore.Seed()

	log.Printf("Seeding done")
	return error(nil)
}
