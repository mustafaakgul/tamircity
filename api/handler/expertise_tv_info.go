package handler

import (
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/gin-gonic/gin"
)

type expertiseTvInfoHandler struct {
	expertiseTvInfoService service.ExpertiseTvInfoService
}

type ExpertiseTvInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertiseTvInfoHandler(expertiseTvInfoService service.ExpertiseTvInfoService) ExpertiseTvInfoHandler {
	return &expertiseTvInfoHandler{
		expertiseTvInfoService: expertiseTvInfoService,
	}
}

func (e *expertiseTvInfoHandler) Create(ctx *gin.Context) {
	// TODO: Implement
}
