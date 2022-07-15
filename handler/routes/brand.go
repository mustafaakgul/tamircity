package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func BrandRouter(router *gin.Engine, brandHandler api.BrandHandler) {
	route := router.Group("api/v1/brand")
	{
		route.GET("/brand", brandHandler.GetAll)
		route.GET("/:deviceTypeId", brandHandler.GetAllByDeviceTypeId)
	}
}
