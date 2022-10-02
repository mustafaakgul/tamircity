package tech_service

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func TechnicalServiceShiftRouter(router *gin.Engine, technicalServiceShiftHandler handler.TechnicalServiceShiftHandler) {
	route := router.Group("api/v1/tech/technical-service-shifts")
	{
		route.GET("/query", technicalServiceShiftHandler.GetByTechnicalServiceId)
		route.POST("", technicalServiceShiftHandler.Create)
	}
}
