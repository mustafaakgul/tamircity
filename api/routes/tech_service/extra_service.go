package tech_service

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExtraServiceRouter(router *gin.Engine, extraServiceHandler handler.ExtraServiceHandler) {
	route := router.Group("api/v1/tech/extra-services")
	{
		route.GET("", extraServiceHandler.GetAll)
	}
}
