package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type modelHandler struct {
	modelService service.ModelService
}

type ModelHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByBrandIdDeviceTypeId(ctx *gin.Context)
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

func (m *modelHandler) GetAllByBrandIdDeviceTypeId(ctx *gin.Context) {

	brandId, err := strconv.Atoi(ctx.Query("brand_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	deviceTypeId, err := strconv.Atoi(ctx.Query("device_type_id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	models, err := m.modelService.FindByBrandIdDeviceTypeId(brandId, deviceTypeId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	response := utils.HandleResponseModel(true, "", nil, models)
	ctx.JSON(http.StatusOK, response)
}
