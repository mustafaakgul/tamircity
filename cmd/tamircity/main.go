package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mustafakocatepe/Tamircity/api/handler"
	"github.com/mustafakocatepe/Tamircity/api/routes"
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	postgres "github.com/mustafakocatepe/Tamircity/pkg/store/shared/db"
	"log"
)

func main() {
	// Set enviroment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	// Connection DB and migrations
	if !db.Migrator().HasTable(&dbModels.Brand{}) {
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
		)
		// Adding Seed data
		technicalServiceStore := repositories.NewTechnicalServiceStore(db)
		/*serviceTypeStore := repositories.NewServiceTypeStore(db)
		extraServiceStore := repositories.NewExtraServiceStore(db)
		brandStore := repositories.NewBrandStore(db)
		modelStore := repositories.NewModelStore(db)
		fixTypeStore := repositories.NewFixTypeStore(db)
		deviceTypeStore := repositories.NewDeviceTypeStore(db)*/
		technicalServiceStore.Seed()
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

	// TODO: Migrating
	/*db.AutoMigrate(
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
	)*/
	log.Println("Migrations done")

	// Store
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	fixTypeStore := repositories.NewFixTypeStore(db)
	deviceTypeStore := repositories.NewDeviceTypeStore(db)

	// TODO: Adding Seed data
	//technicalServiceStore.Seed()

	// Clients
	// This one need to be integrated systems

	// Service
	technicalServiceService := service.NewTechnicalServiceService(technicalServiceStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	extraServiceService := service.ExtraServiceService(extraServiceStore)
	brandService := service.NewBrandService(brandStore)
	modelService := service.NewModelService(modelStore)
	fixTypeService := service.NewFixTypeService(fixTypeStore)
	deviceTypeService := service.NewDeviceTypeService(deviceTypeStore)

	// Handler
	technicalServiceHandler := handler.NewTechnicalServiceHandler(technicalServiceService)
	serviceTypeHandler := handler.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := handler.NewExtraServiceHandler(extraServiceService)
	brandHandler := handler.NewBrandHandler(brandService)
	modelHandler := handler.NewModelHandler(modelService)
	fixTypeHandler := handler.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := handler.NewDeviceTypeHandler(deviceTypeService)

	// gin server
	router := gin.Default()
	router.Use(gin.Logger())

	// CORS Policy
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	router.Use(cors.New(corsConfig))

	// Routes
	routes.BrandRouter(router, brandHandler)
	routes.DeviceTypeRouter(router, deviceTypeHandler)
	routes.ExtraServiceRouter(router, extraServiceHandler)
	routes.FixTypeRouter(router, fixTypeHandler)
	routes.ModelRouter(router, modelHandler)
	routes.ServiceTypeRouter(router, serviceTypeHandler)
	routes.TechnicalServiceRouter(router, technicalServiceHandler)

	router.Run(":8080")
}
