package tech_service

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func ReservationRouter(router *gin.Engine, reservationHandler tech_service.ReservationHandler) {
	route := router.Group("api/v1/tech/reservations")
	{
		route.POST("", reservationHandler.Create)
		route.PATCH("/query", reservationHandler.UpdateReservationStatus)
		route.GET("/pending-and-completed-count", reservationHandler.GetPendingAndCompletedReservationCount)
		route.GET("/:id", reservationHandler.FindByID)
		route.GET("/pending", reservationHandler.GetPendingList)
		route.GET("/completed", reservationHandler.GetCompletedList)
		route.GET("/cancelled", reservationHandler.GetCancelledList)
		route.GET("/approved", reservationHandler.GetApprovedList)
		route.GET("/approved/query", reservationHandler.GetApprovedListByTechnicalServiceIdAndDatetime)
		route.GET("/change-operation-status/query", reservationHandler.ChangeOperationStatus)
	}
}
