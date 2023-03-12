package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func PaymentInfoRouter(router *gin.Engine, paymentInfoHandler handler.PaymentInfoHandler) {
	route := router.Group("api/v1/payments")
	{
		route.POST("", paymentInfoHandler.Create)
		route.GET("/query", paymentInfoHandler.GetAllByExpertiseServiceId)
		route.GET("/:id", paymentInfoHandler.GetById)
	}
}
