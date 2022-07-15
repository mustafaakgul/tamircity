package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/handler/api"
)

func ServiceTypeRouter(router *gin.Engine, serviceTypeHandler api.ServiceTypeHandler) {
	route := router.Group("api/v1/service-type")
	{
		route.GET("/", serviceTypeHandler.GetAll)
	}
}
