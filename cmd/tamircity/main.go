package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/store/repositories"
	"log"

	"github.com/joho/godotenv"
	dbModels "github.com/mustafakocatepe/Tamircity/pkg/models/db"
	postgres "github.com/mustafakocatepe/Tamircity/pkg/store/shared/db"
)

func main() {
	//Set enviroment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init", err)
	}
	log.Println("Postgres connected")

	db.AutoMigrate(&dbModels.Brand{}, &dbModels.Model{}, &dbModels.FixType{}, &dbModels.DeviceType{})
	log.Println("Migrations done")

	//Store
	technicalServiceStore := repositories.NewTechnicalServiceStore(db)
	serviceTypeStore := repositories.NewServiceTypeStore(db)
	extraServiceStore := repositories.NewExtraServiceStore(db)
	brandStore := repositories.NewBrandRepository(db)
	modelStore := repositories.NewModelRepository(db)
	fixTypeStore := repositories.NewFixTypeRepository(db)
	deviceTypeStore := repositories.NewDeviceTypeRepository(db)

	//Service
	technicalServiceService := service.NewTechnicalServiceService(technicalServiceStore)
	serviceTypeService := service.NewServiceTypeService(serviceTypeStore)
	extraServiceService := service.ExtraServiceService(extraServiceStore)
	brandService := service.NewBrandService(brandStore)
	modelService := service.NewModelService(modelStore)
	fixTypeService := service.NewFixTypeService(fixTypeStore)
	deviceTypeService := service.NewDeviceTypeService(deviceTypeStore)

	//Handler
	serviceTypeHandler := api.NewServiceTypeHandler(serviceTypeService)
	extraServiceHandler := api.NewExtraServiceHandler(extraServiceService)
	brandHandler := api.NewBrandHandler(brandService)
	modelHandler := api.NewModelHandler(modelService)
	fixTypeHandler := api.NewFixTypeHandler(fixTypeService)
	deviceTypeHandler := api.NewDeviceTypeHandler(deviceTypeService)

	//gin server
	router := gin.Default()
	router.Use(gin.Logger())

	route := router.Group("api/v1")
	{
		route.GET("/technical-service", serviceTypeHandler.GetAll)
		route.GET("/service_type", serviceTypeHandler.GetAll)
		route.GET("/extra_service", extraServiceHandler.GetAll)
		route.GET("/brand", brandHandler.GetAll)
		route.GET("/model", modelHandler.GetAll)
		route.GET("/fix_type", fixTypeHandler.GetAll)
		route.GET("/device_type", deviceTypeHandler.GetAll)
	}

	router.Run(":8080")
}
