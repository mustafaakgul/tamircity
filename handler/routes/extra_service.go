package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func ExtraServiceRouter(router *gin.Engine, extraServiceHandler api.ExtraServiceHandler) {
	route := router.Group("api/v1/extra-service")
	{
		route.GET("/", extraServiceHandler.GetAll)
	}
}
