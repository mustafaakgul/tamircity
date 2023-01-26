package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertisePcInfoRouter(router *gin.Engine, expertisePcInfoHandler handler.ExpertisePcInfoHandler) {
	route := router.Group("api/v1/expertise_pc_infos")
	{
		route.POST("", expertisePcInfoHandler.Create)
		route.GET("/:id", expertisePcInfoHandler.GetByID)
	}
}
