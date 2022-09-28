package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func TechnicalServiceCandidateRouter(router *gin.Engine, technicalServiceCandidateHandler handler.TechnicalServiceCandidateHandler) {
	technicalServiceCandidateRoute := router.Group("api/v1/technical-services-candidate")
	{
		technicalServiceCandidateRoute.POST("/apply", technicalServiceCandidateHandler.Create)
		technicalServiceCandidateRoute.PUT("/:id", technicalServiceCandidateHandler.Update)
		technicalServiceCandidateRoute.GET("/:id", technicalServiceCandidateHandler.FindByID)
		technicalServiceCandidateRoute.POST("/send-agreement", technicalServiceCandidateHandler.PrepareAndSendAgreement)
		technicalServiceCandidateRoute.GET("/change-status/query", technicalServiceCandidateHandler.ChangeStatusAndCreateTechnicalService)
	}
}
