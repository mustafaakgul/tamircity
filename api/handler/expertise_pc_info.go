package handler

import (
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/gin-gonic/gin"
)

type expertisePcInfoHandler struct {
	expertisePcInfoService service.ExpertisePcInfoService
}

type ExpertisePcInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertisePcInfoHandler(expertisePcInfoService service.ExpertisePcInfoService) ExpertisePcInfoHandler {
	return &expertisePcInfoHandler{
		expertisePcInfoService: expertisePcInfoService,
	}
}

func (e *expertisePcInfoHandler) Create(ctx *gin.Context) {
	// TODO: Implement
}
