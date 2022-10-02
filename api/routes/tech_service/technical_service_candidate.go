package tech_service

import (
	"github.com/anthophora/tamircity/api/handler/tech_service"
	"github.com/gin-gonic/gin"
)

func TechnicalServiceCandidateRouter(router *gin.Engine, technicalServiceCandidateHandler tech_service.TechnicalServiceCandidateHandler) {
	technicalServiceCandidateRoute := router.Group("api/v1/tech/technical-services-candidate")
	{
		technicalServiceCandidateRoute.POST("/apply", technicalServiceCandidateHandler.Create)
		technicalServiceCandidateRoute.PUT("/:id", technicalServiceCandidateHandler.Update)
		technicalServiceCandidateRoute.GET("/:id", technicalServiceCandidateHandler.FindByID)
		technicalServiceCandidateRoute.POST("/send-agreement", technicalServiceCandidateHandler.PrepareAndSendAgreement)
		technicalServiceCandidateRoute.GET("/change-status/query", technicalServiceCandidateHandler.ChangeStatusAndCreateTechnicalService)
	}
}
