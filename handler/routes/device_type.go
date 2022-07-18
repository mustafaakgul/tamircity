package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func DeviceTypeRouter(router *gin.Engine, deviceTypeHandler api.DeviceTypeHandler) {
	route := router.Group("api/v1/device-type")
	{
		//route.GET("/", deviceTypeHandler.GetAll)
		route.GET("/", deviceTypeHandler.GetAllByActive)
	}
}
