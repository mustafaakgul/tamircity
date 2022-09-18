package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func FixTypeRouter(router *gin.Engine, fixTypeHandler handler.FixTypeHandler) {
	fixTypeRoute := router.Group("api/v1/fix-types")
	{
		fixTypeRoute.GET("", fixTypeHandler.GetAll)
		fixTypeRoute.GET("/query", fixTypeHandler.GetAllByDevicetypeId)
	}
}
