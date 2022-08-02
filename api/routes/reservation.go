package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func ReservationRouter(router *gin.Engine, reservationHandler handler.ReservationHandler) {
	route := router.Group("api/v1/reservation")
	{
		route.POST("/", reservationHandler.Create)
	}
}
