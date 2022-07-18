package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func FixTypeRouter(router *gin.Engine, fixTypeHandler api.FixTypeHandler) {
	fixTypeRoute := router.Group("api/v1/fix-type")
	{
		//fixTypeRoute.GET("/", fixTypeHandler.GetAll)
		fixTypeRoute.GET("/", fixTypeHandler.GetAllByDevicetypeId)
	}
}
