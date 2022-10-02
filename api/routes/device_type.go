package routes

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func DeviceTypeRouter(router *gin.Engine, deviceTypeHandler tech_service.DeviceTypeHandler) {
	route := router.Group("api/v1/device-types")
	{
		route.GET("", deviceTypeHandler.GetAll)
		route.GET("/query", deviceTypeHandler.GetAllByActive)
		route.POST("", deviceTypeHandler.Create)
		route.GET("/:id", deviceTypeHandler.GetById)
	}
}
