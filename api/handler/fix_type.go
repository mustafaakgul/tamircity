package handler

import (
	"github.com/anthophora/tamircity/pkg/service/tech_service"
	"net/http"
	"strconv"

	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type fixTypeHandler struct {
	fixTypeService tech_service.FixTypeService
}

type FixTypeHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByDevicetypeId(ctx *gin.Context)
}

func NewFixTypeHandler(fixTypeService tech_service.FixTypeService) FixTypeHandler {
	return &fixTypeHandler{
		fixTypeService: fixTypeService,
	}
}

func (f *fixTypeHandler) GetAll(ctx *gin.Context) {
	fixTypes, err := f.fixTypeService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, fixTypes)
	ctx.JSON(http.StatusOK, response)
}

func (f *fixTypeHandler) GetAllByDevicetypeId(ctx *gin.Context) {

	deviceTypeId, err := strconv.Atoi(ctx.Query("device_type_Id"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	fixTypes, err := f.fixTypeService.FindByDeviceTypeId(deviceTypeId)

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, fixTypes)
	ctx.JSON(http.StatusOK, response)
}
