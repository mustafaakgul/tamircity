package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func ModelRouter(router *gin.Engine, modelHandler api.ModelHandler) {
	route := router.Group("api/v1/model")
	{
		//route.GET("/", modelHandler.GetAll)
		route.GET("/", modelHandler.GetAllByBrandIdDeviceTypeId)
	}
}
