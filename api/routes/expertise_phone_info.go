package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertisePhoneInfoRouter(router *gin.Engine, expertisePhoneInfoHandler handler.ExpertisePhoneInfoHandler) {
	route := router.Group("api/v1/expertise_phone_infos")
	{
		route.POST("", expertisePhoneInfoHandler.Create)
	}
}
