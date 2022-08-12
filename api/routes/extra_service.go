package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func ExtraServiceRouter(router *gin.Engine, extraServiceHandler handler.ExtraServiceHandler) {
	route := router.Group("api/v1/extra-services")
	{
		route.GET("/", extraServiceHandler.GetAll)
	}
}
