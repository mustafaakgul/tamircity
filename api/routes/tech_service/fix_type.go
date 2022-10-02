package tech_service

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func FixTypeRouter(router *gin.Engine, fixTypeHandler tech_service.FixTypeHandler) {
	fixTypeRoute := router.Group("api/v1/tech/fix-types")
	{
		fixTypeRoute.GET("", fixTypeHandler.GetAll)
		fixTypeRoute.GET("/query", fixTypeHandler.GetAllByDevicetypeId)
	}
}
