package tech_service

import (
	"github.com/anthophora/tamircity/pkg/service/tech_service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type extraServiceHandler struct {
	extraServiceService tech_service.ExtraServiceService
}

type ExtraServiceHandler interface {
	GetAll(ctx *gin.Context)
}

func NewExtraServiceHandler(extraServiceService tech_service.ExtraServiceService) ExtraServiceHandler {
	return &extraServiceHandler{
		extraServiceService: extraServiceService,
	}
}

func (e *extraServiceHandler) GetAll(ctx *gin.Context) {
	extraServices, err := e.extraServiceService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, extraServices)
	ctx.JSON(http.StatusOK, response)
}
