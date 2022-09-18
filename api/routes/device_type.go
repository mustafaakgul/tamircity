package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func DeviceTypeRouter(router *gin.Engine, deviceTypeHandler handler.DeviceTypeHandler) {
	route := router.Group("api/v1/device-types")
	{
		route.GET("", deviceTypeHandler.GetAll)
		route.GET("/query", deviceTypeHandler.GetAllByActive)
		route.POST("", deviceTypeHandler.Create)
		route.GET("/:id", deviceTypeHandler.GetById)
	}
}
