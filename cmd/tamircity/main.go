package main

import (
	"github.com/anthophora/tamircity/api/handler"
	tech_service4 "github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/anthophora/tamircity/api/routes"
	tech_service2 "github.com/anthophora/tamircity/api/routes/tech_service"
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/anthophora/tamircity/pkg/service"
	tech_service3 "github.com/anthophora/tamircity/pkg/service/tech_service"
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
	technicalServiceService := tech_service3.NewTechnicalServiceService(technicalServiceStore)
	technicalServiceCandidateService := tech_service3.NewTechnicalServiceCandidateService(technicalServiceCandidateStore)
	serviceTypeService := tech_service3.NewServiceTypeService(serviceTypeStore)
	extraServiceService := tech_service3.NewExtraServiceService(extraServiceStore)
	brandService := tech_service3.NewBrandService(brandStore)
	modelService := tech_service3.NewModelService(modelStore)
	fixTypeService := tech_service3.NewFixTypeService(fixTypeStore)
	deviceTypeService := tech_service3.NewDeviceTypeService(deviceTypeStore)
	reservationService := tech_service3.NewReservationService(reservationStore)
	userService := service.NewUserService(userStore)
	technicalServiceShiftService := tech_service3.NewTechnicalServiceShiftService(technicalServiceShiftStore)

	// Handler
	technicalServiceHandler := tech_service4.NewTechnicalServiceHandler(technicalServiceService)
	technicalServiceCandidateHandler := tech_service4.NewTechnicalServiceCandidateHandler(technicalServiceCandidateService)
	serviceTypeHandler := tech_service4.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := tech_service4.NewExtraServiceHandler(extraServiceService)
	brandHandler := tech_service4.NewBrandHandler(brandService)
	modelHandler := tech_service4.NewModelHandler(modelService)
	fixTypeHandler := tech_service4.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := tech_service4.NewDeviceTypeHandler(deviceTypeService)
	reservationHandler := tech_service4.NewReservationHandler(reservationService)
	userHandler := handler.NewUserHandler(userService)
	technicalServiceShiftHandler := tech_service4.NewTechnicalServiceShiftHandler(technicalServiceShiftService)

	// Gin Server
	router := gin.Default()
	router.Use(gin.Logger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	router.Use(cors.New(corsConfig))

	// Routes
	tech_service2.BrandRouter(router, brandHandler)
	tech_service2.DeviceTypeRouter(router, deviceTypeHandler)
	tech_service2.ExtraServiceRouter(router, extraServiceHandler)
	tech_service2.FixTypeRouter(router, fixTypeHandler)
	tech_service2.ModelRouter(router, modelHandler)
	tech_service2.ServiceTypeRouter(router, serviceTypeHandler)
	tech_service2.TechnicalServiceRouter(router, technicalServiceHandler)
	tech_service2.TechnicalServiceCandidateRouter(router, technicalServiceCandidateHandler)
	tech_service2.ReservationRouter(router, reservationHandler)
	routes.UserRouter(router, userHandler)
	tech_service2.TechnicalServiceShiftRouter(router, technicalServiceShiftHandler)

	port := os.Getenv("API_PORT")
	router.Run(":" + port)
}
