package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mustafakocatepe/Tamircity/handler/api"
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
	technicalServiceHandler := api.NewTechnicalServiceHandler(technicalServiceService)
	serviceTypeHandler := api.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := api.NewExtraServiceHandler(extraServiceService)
	brandHandler := api.NewBrandHandler(brandService)
	modelHandler := api.NewModelHandler(modelService)
	fixTypeHandler := api.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := api.NewDeviceTypeHandler(deviceTypeService)

	// gin server
	router := gin.Default()
	router.Use(gin.Logger())

	// CORS Policy
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	router.Use(cors.New(corsConfig))

	route := router.Group("api/v1")
	{
		// Get All Models APIs
		route.GET("/technical-service", technicalServiceHandler.GetAll)
		route.GET("/service-type", serviceTypeHandler.GetAll)
		route.GET("/extra-service", extraServiceHandler.GetAll)
		route.GET("/brand", brandHandler.GetAll)
		route.GET("/model", modelHandler.GetAll)
		route.GET("/fix-type", fixTypeHandler.GetAll)
		route.GET("/device-type", deviceTypeHandler.GetAll)

		// Create Model APIs
		route.POST("/technical-service", technicalServiceHandler.Create)

		// Get Model APIs
		route.GET("/technical-service/:id", technicalServiceHandler.Get)

		// Update Model APIs
		route.PUT("/technical-service/:id", technicalServiceHandler.Get)

		// Delete Model APIs
		route.DELETE("/technical-service/:id", technicalServiceHandler.Delete)
	}

	router.Run(":8080")
}
