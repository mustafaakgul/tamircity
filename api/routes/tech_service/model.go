package tech_service

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ModelRouter(router *gin.Engine, modelHandler handler.ModelHandler) {
	route := router.Group("api/v1/tech/models")
	{
		route.GET("", modelHandler.GetAll)
		route.GET("/query", modelHandler.GetAllByBrandIdDeviceTypeId)
		route.POST("", modelHandler.Create)
	}
}
