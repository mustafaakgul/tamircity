package postgres

import (
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	// Connection DB and migrations
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
		log.Println("Migrations done")
	}
	return error(nil)
}
