package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/pkg/service"
	"github.com/mustafakocatepe/Tamircity/pkg/utils"
)

type deviceTypeHandler struct {
	deviceTypeService service.DeviceTypeService
}

type DeviceTypeHandler interface {
	GetAll(ctx *gin.Context)
}

func NewDeviceTypeHandler(deviceTypeService service.DeviceTypeService) DeviceTypeHandler {
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
