package main

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/anthophora/tamircity/api/routes"
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/store/repositories"
	"github.com/anthophora/tamircity/pkg/store/repositories/tech_service"
	postgres "github.com/anthophora/tamircity/pkg/store/shared/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Logger initiliazation
	middleware.SentryLogger()

	// Set Enviroment Variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connection to Postgres
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	// Migrations Mechanism
	postgres.Migrate(db)

	// Store
	technicalServiceStore := tech_service.NewTechnicalServiceStore(db)
	technicalServiceCandidateStore := tech_service.NewTechnicalServiceCandidateStore(db)
	serviceTypeStore := tech_service.NewServiceTypeStore(db)
	extraServiceStore := tech_service.NewExtraServiceStore(db)
	brandStore := tech_service.NewBrandStore(db)
	modelStore := tech_service.NewModelStore(db)
	reservationStore := tech_service.NewReservationStore(db)
	fixTypeStore := tech_service.NewFixTypeStore(db)
	deviceTypeStore := tech_service.NewDeviceTypeStore(db)
	userStore := repositories.NewUserStore(db)
	technicalServiceShiftStore := tech_service.NewTechnicalServiceShiftStore(db)

	// Clients
	// This one need to be integrated systems

	// Service
	technicalServiceService := service.NewTechnicalServiceService(technicalServiceStore)
	technicalServiceCandidateService := service.NewTechnicalServiceCandidateService(technicalServiceCandidateStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	extraServiceService := service.NewExtraServiceService(extraServiceStore)
	brandService := service.NewBrandService(brandStore)
	modelService := service.NewModelService(modelStore)
	fixTypeService := service.NewFixTypeService(fixTypeStore)
	deviceTypeService := service.NewDeviceTypeService(deviceTypeStore)
	reservationService := service.NewReservationService(reservationStore)
	userService := service.NewUserService(userStore)
	technicalServiceShiftService := service.NewTechnicalServiceShiftService(technicalServiceShiftStore)

	// Handler
	technicalServiceHandler := handler.NewTechnicalServiceHandler(technicalServiceService)
	technicalServiceCandidateHandler := handler.NewTechnicalServiceCandidateHandler(technicalServiceCandidateService)
	serviceTypeHandler := handler.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := handler.NewExtraServiceHandler(extraServiceService)
	brandHandler := handler.NewBrandHandler(brandService)
	modelHandler := handler.NewModelHandler(modelService)
	fixTypeHandler := handler.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := handler.NewDeviceTypeHandler(deviceTypeService)
	reservationHandler := handler.NewReservationHandler(reservationService)
	userHandler := handler.NewUserHandler(userService)
	technicalServiceShiftHandler := handler.NewTechnicalServiceShiftHandler(technicalServiceShiftService)

	// Gin Server
	router := gin.Default()
	router.Use(gin.Logger())

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
	routes.TechnicalServiceCandidateRouter(router, technicalServiceCandidateHandler)
	routes.ReservationRouter(router, reservationHandler)
	routes.UserRouter(router, userHandler)
	routes.TechnicalServiceShiftRouter(router, technicalServiceShiftHandler)

	port := os.Getenv("API_PORT")
	router.Run(":" + port)
}
