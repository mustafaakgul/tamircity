package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func TechnicalServiceRouter(router *gin.Engine, technicalServiceHandler handler.TechnicalServiceHandler) {
	technicalServiceRoute := router.Group("api/v1/technical-service")
	{
		technicalServiceRoute.GET("/", technicalServiceHandler.GetAll)
		technicalServiceRoute.GET("/query", technicalServiceHandler.GetAllByFilter)
		technicalServiceRoute.POST("/", technicalServiceHandler.Create)
		technicalServiceRoute.GET("/:id", technicalServiceHandler.Get)
		technicalServiceRoute.PUT("/:id", technicalServiceHandler.Get)
		technicalServiceRoute.DELETE("/:id", technicalServiceHandler.Delete)
	}
}
