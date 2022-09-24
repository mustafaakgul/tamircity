package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ReservationRouter(router *gin.Engine, reservationHandler handler.ReservationHandler) {
	route := router.Group("api/v1/reservations")
	{
		route.POST("", reservationHandler.Create)
		route.PATCH("/query", reservationHandler.UpdateReservationStatus)
		route.GET("/pending-and-completed-count", reservationHandler.GetPendingAndCompletedReservationCount)
		route.GET("/pending", reservationHandler.GetPendingList)
		route.GET("/completed", reservationHandler.GetCompletedList)
		route.GET("/cancelled", reservationHandler.GetCancelledList)
		route.GET("/approved", reservationHandler.GetApprovedList)
		route.GET("/approved/query", reservationHandler.GetApprovedListByTechnicalServiceIdAndDatetime)
	}
}
