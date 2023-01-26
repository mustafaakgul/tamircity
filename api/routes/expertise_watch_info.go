package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseWatchInfoRouter(router *gin.Engine, expertiseWatchInfoHandler handler.ExpertiseWatchInfoHandler) {
	route := router.Group("api/v1/expertise_watch_infos")
	{
		route.POST("", expertiseWatchInfoHandler.Create)
		route.GET("/:id", expertiseWatchInfoHandler.GetByID)
	}
}
