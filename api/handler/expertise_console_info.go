package handler

import (
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/gin-gonic/gin"
)

type expertiseConsoleInfoHandler struct {
	expertiseConsoleInfoService service.ExpertiseConsoleInfoService
}

type ExpertiseConsoleInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertiseConsoleInfoHandler(expertiseConsoleInfoService service.ExpertiseConsoleInfoService) ExpertiseConsoleInfoHandler {
	return &expertiseConsoleInfoHandler{
		expertiseConsoleInfoService: expertiseConsoleInfoService,
	}
}

func (e *expertiseConsoleInfoHandler) Create(ctx *gin.Context) {
	// TODO: Implement
}
