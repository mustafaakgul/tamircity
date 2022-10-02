package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseServiceShiftRouter(router *gin.Engine, expertiseServiceShiftHandler handler.ExpertiseServiceShiftHandler) {
	route := router.Group("api/v1/expertise-service-shifts")
	{
		route.GET("/query", expertiseServiceShiftHandler.GetByExpertiseServiceId)
		route.POST("", expertiseServiceShiftHandler.Create)
	}
}
