package main

import (
	"log"
	"os"

	"github.com/anthophora/tamircity/api/handler"
	"github.com/anthophora/tamircity/api/routes"
	"github.com/anthophora/tamircity/pkg/middleware"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/store/repositories"
	postgres "github.com/anthophora/tamircity/pkg/store/shared/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Logger Initiliazation
	middleware.SentryLogger()

	// Set Enviroment Variables
	err := godotenv.Load("../../.env")
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
	expertiseServiceStore := repositories.NewExpertiseServiceStore(db)
	expertiseServiceCandidateStore := repositories.NewExpertiseServiceCandidateStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	brandStore := repositories.NewBrandStore(db)
	modelStore := repositories.NewModelStore(db)
	reservationStore := repositories.NewReservationStore(db)
	deviceTypeStore := repositories.NewDeviceTypeStore(db)
	userStore := repositories.NewUserStore(db)
	expertiseServiceShiftStore := repositories.NewExpertiseServiceShiftStore(db)

	expertiseConsoleInfoStore := repositories.NewExpertiseConsoleInfoStore(db)
	expertisePcInfoStoreStore := repositories.NewExpertisePcInfoStoreStore(db)
	expertisePhoneInfoStore := repositories.NewExpertisePhoneInfoStore(db)
	expertiseTvInfoStore := repositories.NewExpertiseTvInfoStore(db)
	expertiseWatchInfoStoreStore := repositories.NewExpertiseWatchInfoStoreStore(db)

	// Clients
	// This one need to be integrated systems

	// Service
	expertiseServiceService := service.NewExpertiseServiceService(expertiseServiceStore)
	expertiseServiceCandidateService := service.NewExpertiseServiceCandidateService(expertiseServiceCandidateStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	brandService := service.NewBrandService(brandStore)
	modelService := service.NewModelService(modelStore)
	deviceTypeService := service.NewDeviceTypeService(deviceTypeStore)
	reservationService := service.NewReservationService(reservationStore)
	userService := service.NewUserService(userStore)
	expertiseServiceShiftService := service.NewExpertiseServiceShiftService(expertiseServiceShiftStore)

	expertiseConsoleInfoService := service.NewExpertiseConsoleInfoService(expertiseConsoleInfoStore)
	expertisePCInfoService := service.NewExpertisePCInfoService(expertisePcInfoStoreStore)
	expertisePhoneInfoService := service.NewExpertisePhoneInfoService(expertisePhoneInfoStore)
	expertiseTvInfoService := service.NewExpertiseTvInfoService(expertiseTvInfoStore)
	expertiseWatchInfoService := service.NewExpertiseWatchInfoService(expertiseWatchInfoStoreStore)

	// Handler
	expertiseServiceHandler := handler.NewExpertiseServiceHandler(expertiseServiceService)
	expertiseServiceCandidateHandler := handler.NewExpertiseServiceCandidateHandler(expertiseServiceCandidateService)
	serviceTypeHandler := handler.NewServiceTypeHandler(serviceTypeService)
	brandHandler := handler.NewBrandHandler(brandService)
	modelHandler := handler.NewModelHandler(modelService)
	deviceTypeHandler := handler.NewDeviceTypeHandler(deviceTypeService)
	reservationHandler := handler.NewReservationHandler(reservationService)
	userHandler := handler.NewUserHandler(userService)
	expertiseServiceShiftHandler := handler.NewExpertiseServiceShiftHandler(expertiseServiceShiftService)

	expertiseConsoleInfoHandler := handler.NewExpertiseConsoleInfoHandler(expertiseConsoleInfoService)
	expertisePcInfoHandler := handler.NewExpertisePcInfoHandler(expertisePCInfoService)
	expertisePhoneInfoHandler := handler.NewExpertisePhoneInfoHandler(expertisePhoneInfoService)
	expertiseTvInfoHandler := handler.NewExpertiseTvInfoHandler(expertiseTvInfoService)
	expertiseWatchInfoHandler := handler.NewExpertiseWatchInfoHandler(expertiseWatchInfoService)

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
	routes.ModelRouter(router, modelHandler)
	routes.ServiceTypeRouter(router, serviceTypeHandler)
	routes.ExpertiseServiceRouter(router, expertiseServiceHandler)
	routes.ExpertiseServiceCandidateRouter(router, expertiseServiceCandidateHandler)
	routes.ReservationRouter(router, reservationHandler)
	routes.UserRouter(router, userHandler)
	routes.ExpertiseServiceShiftRouter(router, expertiseServiceShiftHandler)

	routes.ExpertiseConsoleInfoRouter(router, expertiseConsoleInfoHandler)
	routes.ExpertisePcInfoRouter(router, expertisePcInfoHandler)
	routes.ExpertisePhoneInfoRouter(router, expertisePhoneInfoHandler)
	routes.ExpertiseTvInfoRouter(router, expertiseTvInfoHandler)
	routes.ExpertiseWatchInfoRouter(router, expertiseWatchInfoHandler)

	port := os.Getenv("API_PORT")
	router.Run(":" + port)
}
