package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
	"net/http"
)

type extraServiceHandler struct {
	extraServiceService service.ExtraServiceService
}

type ExtraServiceHandler interface {
	GetAll(ctx *gin.Context)
}

func NewExtraServiceHandler(extraServiceService service.ExtraServiceService) ExtraServiceHandler {
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
