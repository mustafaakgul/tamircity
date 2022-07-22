package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func BrandRouter(router *gin.Engine, brandHandler handler.BrandHandler) {
	route := router.Group("api/v1/brand")
	{
		route.GET("/", brandHandler.GetAll)
		route.GET("/query", brandHandler.GetAllByDeviceTypeId)
	}
}
