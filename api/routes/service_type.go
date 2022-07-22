package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func ServiceTypeRouter(router *gin.Engine, serviceTypeHandler handler.ServiceTypeHandler) {
	route := router.Group("api/v1/service-type")
	{
		route.GET("/", serviceTypeHandler.GetAll)
	}
}
