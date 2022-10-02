package tech_service

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func BrandRouter(router *gin.Engine, brandHandler tech_service.BrandHandler) {
	route := router.Group("api/v1/tech/brands")
	{
		route.GET("", brandHandler.GetAll)
		route.GET("/query", brandHandler.GetAllByDeviceTypeId)
		route.POST("", brandHandler.Create)
	}
}
