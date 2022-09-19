package postgres

import (
	"log"

	"github.com/anthophora/tamircity/pkg/store/repositories"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	// Adding Seed data
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	fixTypeStore := repositories.NewFixTypeStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)
	/*deviceTypeStore := repositories.NewDeviceTypeStore(db)
	 */
	technicalServiceStore.Seed()
	brandStore.Seed()
	modelStore.Seed()
	fixTypeStore.Seed()
	serviceTypeStore.Seed()
	extraServiceStore.Seed()

	log.Printf("Seeding done")

	return error(nil)
}
