package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db/tech_service"
	tech_service3 "github.com/anthophora/tamircity/pkg/models/web/tech_service"
	tech_service2 "github.com/anthophora/tamircity/pkg/service/tech_service"
	"net/http"
	"strconv"

	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type deviceTypeHandler struct {
	deviceTypeService tech_service2.DeviceTypeService
}

type DeviceTypeHandler interface {
	GetAll(ctx *gin.Context)
	GetAllByActive(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

func NewDeviceTypeHandler(deviceTypeService tech_service2.DeviceTypeService) DeviceTypeHandler {
	return &deviceTypeHandler{
		deviceTypeService: deviceTypeService,
	}
}

func (d *deviceTypeHandler) GetAll(ctx *gin.Context) {
	deviceTypes, err := d.deviceTypeService.FindAll()

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, deviceTypes)
	ctx.JSON(http.StatusOK, response)
}

func (d *deviceTypeHandler) GetAllByActive(ctx *gin.Context) {

	active, err := strconv.Atoi(ctx.Query("active"))
	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}
	var deviceTypes []tech_service3.DeviceTypeResponse
	if active == 1 {
		deviceTypes, err = d.deviceTypeService.FindAllByActive()
	}

	if err != nil {
		responseErr := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, responseErr)
		return
	}

	response := utils.HandleResponseModel(true, "", nil, deviceTypes)
	ctx.JSON(http.StatusOK, response)
}

func (d *deviceTypeHandler) Create(ctx *gin.Context) {
	var deviceType tech_service3.DeviceTypeRequest
	if err := ctx.ShouldBindJSON(&deviceType); err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var deviceTypeModel tech_service.DeviceType
	deviceTypeModel.Name = deviceType.Name
	deviceTypeModel.ShortDescription = deviceType.ShortDescription
	deviceTypeModel.IsActive = deviceType.IsActive

	err := d.deviceTypeService.Create(&deviceTypeModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
	response := utils.HandleResponseModel(true, "", nil, deviceTypeModel)
	ctx.JSON(http.StatusOK, response)
}

func (d *deviceTypeHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	deviceType, _ := d.deviceTypeService.FindByID(id)
	ctx.JSON(http.StatusOK, deviceType)
}
