package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func FixTypeRouter(router *gin.Engine, fixTypeHandler handler.FixTypeHandler) {
	fixTypeRoute := router.Group("api/v1/fix-type")
	{
		fixTypeRoute.GET("/", fixTypeHandler.GetAll)
		fixTypeRoute.GET("/query", fixTypeHandler.GetAllByDevicetypeId)
	}
}
