package handler

import (
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/gin-gonic/gin"
)

type expertisePhoneInfoHandler struct {
	expertisePhoneInfoService service.ExpertisePhoneInfoService
}

type ExpertisePhoneInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertisePhoneInfoHandler(expertisePhoneInfoService service.ExpertisePhoneInfoService) ExpertisePhoneInfoHandler {
	return &expertisePhoneInfoHandler{
		expertisePhoneInfoService: expertisePhoneInfoService,
	}
}

func (e *expertisePhoneInfoHandler) Create(ctx *gin.Context) {
	// TODO: Implement
}
