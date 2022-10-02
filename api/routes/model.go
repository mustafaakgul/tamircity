package routes

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func ModelRouter(router *gin.Engine, modelHandler tech_service.ModelHandler) {
	route := router.Group("api/v1/models")
	{
		route.GET("", modelHandler.GetAll)
		route.GET("/query", modelHandler.GetAllByBrandIdDeviceTypeId)
		route.POST("", modelHandler.Create)
	}
}
