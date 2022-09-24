package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func TechnicalServiceShiftRouter(router *gin.Engine, technicalServiceShiftHandler handler.TechnicalServiceShiftHandler) {
	route := router.Group("api/v1/technical-service-shifts")
	{
		route.GET("/query", technicalServiceShiftHandler.GetByTechnicalServiceId)
		route.POST("", technicalServiceShiftHandler.Create)
	}
}
