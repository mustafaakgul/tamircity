package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type modelHandler struct {
	modelService service.ModelService
}

type ModelHandler interface {
	GetAll(ctx *gin.Context)
}

func NewModelHandler(modelService service.ModelService) ModelHandler {
	return &modelHandler{
		modelService: modelService,
	}
}

func (m *modelHandler) GetAll(ctx *gin.Context) {
	models, err := m.modelService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, models)
	ctx.JSON(http.StatusOK, response)
}
