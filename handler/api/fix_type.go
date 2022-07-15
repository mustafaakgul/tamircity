package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type fixTypeHandler struct {
	fixTypeService service.FixTypeService
}

type FixTypeHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByDevicetypeId(ctx *gin.Context)
}

func NewFixTypeHandler(fixTypeService service.FixTypeService) FixTypeHandler {
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

	deviceTypeId, err := strconv.Atoi(ctx.Param("deviceTypeId"))
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
