package tech_service

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func ServiceTypeRouter(router *gin.Engine, serviceTypeHandler tech_service.ServiceTypeHandler) {
	route := router.Group("api/v1/tech/service-types")
	{
		route.GET("", serviceTypeHandler.GetAll)
	}
}
