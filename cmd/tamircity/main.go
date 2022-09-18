package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mustafakocatepe/Tamircity/api/handler"
	"github.com/mustafakocatepe/Tamircity/api/routes"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	postgres "github.com/mustafakocatepe/Tamircity/pkg/store/shared/db"
	"log"
)

func main() {
	// Set enviroment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	// Connection DB and migrations
	postgres.Migrate(db)

	// TODO: Migrating
	//postgres.Seeder(db)  //TODO
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
	reservationStore := repositories.NewReservationStore(db)
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
	reservationService := service.NewReservationService(reservationStore)

	// Handler
	technicalServiceHandler := handler.NewTechnicalServiceHandler(technicalServiceService)
	serviceTypeHandler := handler.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := handler.NewExtraServiceHandler(extraServiceService)
	brandHandler := handler.NewBrandHandler(brandService)
	modelHandler := handler.NewModelHandler(modelService)
	fixTypeHandler := handler.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := handler.NewDeviceTypeHandler(deviceTypeService)
	reservationHandler := handler.NewReservationHandler(reservationService)

	// CORS Policy
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowOrigins = []string{"*"}
	//corsConfig.AllowCredentials = true
	//corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	//corsConfig.AddAllowHeaders("Access-Control-Allow-Origin", "*")
	/*w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")*/
	//corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	//router.Use(cors.New(corsConfig))

	// gin server
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/line", func(c *gin.Context) {
		// Note: During front-end and back-end separation, attention should be paid to cross-domain issues, so request headers need to be set up.
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		legendData := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
		c.JSON(200, gin.H{
			"legend_data": legendData,
		})
	})

	/*corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8081/", "http://localhost:8081"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "PUT", "PATCH", "POST")
	corsConfig.AddExposeHeaders("Access-Control-Allow-Origin", "*")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Origin", "*")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Origin", "http://localhost:8081/")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Origin", "http://localhost:8081")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Methods: GET")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Methods: OPTIONS")

	router.Use(cors.New(corsConfig))*/

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
	routes.ReservationRouter(router, reservationHandler)

	router.Run(":8000")
}
