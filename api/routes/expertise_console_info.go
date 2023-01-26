package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseConsoleInfoRouter(router *gin.Engine, expertiseConsoleInfoHandler handler.ExpertiseConsoleInfoHandler) {
	route := router.Group("api/v1/expertise_console_infos")
	{
		route.POST("", expertiseConsoleInfoHandler.Create)
		route.GET("/:id", expertiseConsoleInfoHandler.GetByID)
	}
}
