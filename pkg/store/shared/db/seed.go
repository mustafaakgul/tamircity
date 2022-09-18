package postgres

import (
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	"gorm.io/gorm"
	"log"
)

func Seeder(db *gorm.DB) error {
	log.Printf("Seeding started")

	// Adding Seed data
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	fixTypeStore := repositories.NewFixTypeStore(db)
	/*serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	fixTypeStore := repositories.NewFixTypeStore(db)
	deviceTypeStore := repositories.NewDeviceTypeStore(db)*/
	technicalServiceStore.Seed()
	brandStore.Seed()
	modelStore.Seed()
	fixTypeStore.Seed()
	/*roleRepo := role.NewRoleRepository(db)
	roleRepo.Seed()
	userRepo := user.NewUserRepository(db)
	userRepo.Seed()
	statusRepo := status.NewStatusRepository(db)
	statusRepo.Seed()
	users, _ := userRepo.FindAll()
	roles, _ := roleRepo.FindAll()
	userrolemapRepo := userrolemap.NewUserRoleMapRepository(db)
	userrolemapRepo.Seed(users, roles)*/

	log.Printf("Seeding done")

	return error(nil)
}
