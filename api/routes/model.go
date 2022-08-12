package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func ModelRouter(router *gin.Engine, modelHandler handler.ModelHandler) {
	route := router.Group("api/v1/models")
	{
		route.GET("/", modelHandler.GetAll)
		route.GET("/query", modelHandler.GetAllByBrandIdDeviceTypeId)
	}
}
