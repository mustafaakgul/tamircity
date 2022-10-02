package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ServiceTypeRouter(router *gin.Engine, serviceTypeHandler handler.ServiceTypeHandler) {
	route := router.Group("api/v1/service-types")
	{
		route.GET("", serviceTypeHandler.GetAll)
	}
}
