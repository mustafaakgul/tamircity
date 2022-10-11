package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseTvInfoRouter(router *gin.Engine, expertiseTvInfoHandler handler.ExpertiseTvInfoHandler) {
	route := router.Group("api/v1/expertise_tv_infos")
	{
		route.POST("", expertiseTvInfoHandler.Create)
	}
}
