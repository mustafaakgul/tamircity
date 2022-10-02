package routes

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func ExpertiseServiceCandidateRouter(router *gin.Engine, expertiseServiceCandidateHandler handler.ExpertiseServiceCandidateHandler) {
	expertiseServiceCandidateRoute := router.Group("api/v1/expertise-services-candidate")
	{
		expertiseServiceCandidateRoute.POST("/apply", expertiseServiceCandidateHandler.Create)
		expertiseServiceCandidateRoute.PUT("/:id", expertiseServiceCandidateHandler.Update)
		expertiseServiceCandidateRoute.GET("/:id", expertiseServiceCandidateHandler.FindByID)
		expertiseServiceCandidateRoute.POST("/send-agreement", expertiseServiceCandidateHandler.PrepareAndSendAgreement)
		expertiseServiceCandidateRoute.GET("/change-status/query", expertiseServiceCandidateHandler.ChangeStatusAndCreateExpertiseService)
	}
}
