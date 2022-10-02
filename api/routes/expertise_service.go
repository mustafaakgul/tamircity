package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseServiceRouter(router *gin.Engine, expertiseServiceHandler handler.ExpertiseServiceHandler) {
	expertiseServiceRoute := router.Group("api/v1/expertise-services")
	{
		expertiseServiceRoute.GET("", expertiseServiceHandler.GetAll)
		expertiseServiceRoute.GET("/query", expertiseServiceHandler.GetAllByFilter)
		expertiseServiceRoute.POST("", expertiseServiceHandler.Create)
		expertiseServiceRoute.GET("/:id", expertiseServiceHandler.Get)
		expertiseServiceRoute.PUT("/:id", expertiseServiceHandler.Update)
		expertiseServiceRoute.DELETE("/:id", expertiseServiceHandler.Delete)
	}
}
