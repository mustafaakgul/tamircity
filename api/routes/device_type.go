package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
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
