package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"net/http"
	"strconv"

	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type modelHandler struct {
	modelService service.ModelService
}

type ModelHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByBrandIdDeviceTypeId(ctx *gin.Context)
	Create(ctx *gin.Context)
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

func (m *modelHandler) Create(ctx *gin.Context) {
	var model web.ModelRequest
	if err := ctx.ShouldBindJSON(&model); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var modelModel db.Model
	modelModel.Name = model.Name
	modelModel.ShortDescription = model.ShortDescription
	modelModel.IsActive = model.IsActive

	err := m.modelService.Create(&modelModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, modelModel)
	ctx.JSON(http.StatusOK, response)
}
