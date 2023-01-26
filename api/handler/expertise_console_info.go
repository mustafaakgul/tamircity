package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type expertiseConsoleInfoHandler struct {
	expertiseConsoleInfoService service.ExpertiseConsoleInfoService
}

type ExpertiseConsoleInfoHandler interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

func NewExpertiseConsoleInfoHandler(expertiseConsoleInfoService service.ExpertiseConsoleInfoService) ExpertiseConsoleInfoHandler {
	return &expertiseConsoleInfoHandler{
		expertiseConsoleInfoService: expertiseConsoleInfoService,
	}
}

func (e *expertiseConsoleInfoHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := utils.HandleResponseModel(false, "Wrong Params", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	expertiseConsoleInfo, err := e.expertiseConsoleInfoService.GetByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Console Info could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Console Info found successfully", nil, expertiseConsoleInfo)
	ctx.JSON(http.StatusOK, response)
}

func (e *expertiseConsoleInfoHandler) Create(ctx *gin.Context) {
	var expertiseConsoleInfo web.ExpertiseConsoleInfoRequest
	if err := ctx.ShouldBindJSON(&expertiseConsoleInfo); err != nil {
		response := utils.HandleResponseModel(false, "Invalid JSON body", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseConsoleInfoModel db.ExpertiseConsoleInfo
	expertiseConsoleInfoModel.ReservationId = expertiseConsoleInfo.ReservationId
	expertiseConsoleInfoModel.Invoice = expertiseConsoleInfo.Invoice
	expertiseConsoleInfoModel.Box = expertiseConsoleInfo.Box
	expertiseConsoleInfoModel.GuaranteeTerm = expertiseConsoleInfo.GuaranteeTerm
	expertiseConsoleInfoModel.Color = expertiseConsoleInfo.Color
	expertiseConsoleInfoModel.Platform = expertiseConsoleInfo.Platform
	expertiseConsoleInfoModel.StorageCapacity = expertiseConsoleInfo.StorageCapacity
	expertiseConsoleInfoModel.OsType = expertiseConsoleInfo.OsType
	expertiseConsoleInfoModel.CPUModel = expertiseConsoleInfo.CPUModel
	expertiseConsoleInfoModel.FpsValue = expertiseConsoleInfo.FpsValue
	expertiseConsoleInfoModel.HDRSupport = expertiseConsoleInfo.HDRSupport
	expertiseConsoleInfoModel.Ram = expertiseConsoleInfo.Ram
	expertiseConsoleInfoModel.RamType = expertiseConsoleInfo.RamType
	expertiseConsoleInfoModel.GameResolution = expertiseConsoleInfo.GameResolution
	expertiseConsoleInfoModel.VideoResolution = expertiseConsoleInfo.VideoResolution
	expertiseConsoleInfoModel.NetworkType = expertiseConsoleInfo.NetworkType
	expertiseConsoleInfoModel.Ethernet = expertiseConsoleInfo.Ethernet
	expertiseConsoleInfoModel.EthernetSpeed = expertiseConsoleInfo.EthernetSpeed
	expertiseConsoleInfoModel.HDMIStandart = expertiseConsoleInfo.HDMIStandart
	expertiseConsoleInfoModel.USBInput = expertiseConsoleInfo.USBInput
	expertiseConsoleInfoModel.USBInputNumber = expertiseConsoleInfo.USBInputNumber
	expertiseConsoleInfoModel.USBVersion = expertiseConsoleInfo.USBVersion
	expertiseConsoleInfoModel.Bluetooth = expertiseConsoleInfo.Bluetooth
	expertiseConsoleInfoModel.BluetoothVersion = expertiseConsoleInfo.BluetoothVersion
	expertiseConsoleInfoModel.VoiceSupport = expertiseConsoleInfo.VoiceSupport
	expertiseConsoleInfoModel.Controller = expertiseConsoleInfo.Controller
	expertiseConsoleInfoModel.ControllerNumber = expertiseConsoleInfo.ControllerNumber
	expertiseConsoleInfoModel.Game = expertiseConsoleInfo.Game
	expertiseConsoleInfoModel.GameNumber = expertiseConsoleInfo.GameNumber
	expertiseConsoleInfoModel.IsBoxHasProblem = expertiseConsoleInfo.IsBoxHasProblem
	expertiseConsoleInfoModel.IsBoxHasLightProblem = expertiseConsoleInfo.IsBoxHasLightProblem
	expertiseConsoleInfoModel.IsControllerHasLightProblem = expertiseConsoleInfo.IsControllerHasLightProblem
	expertiseConsoleInfoModel.IsControllerHasVibrationProblem = expertiseConsoleInfo.IsControllerHasVibrationProblem
	expertiseConsoleInfoModel.IsControllerHasAnalogProglem = expertiseConsoleInfo.IsControllerHasAnalogProglem
	expertiseConsoleInfoModel.IsControllerHasButtonProblem = expertiseConsoleInfo.IsControllerHasButtonProblem
	expertiseConsoleInfoModel.IsDeviceHasButtonProblem = expertiseConsoleInfo.IsDeviceHasButtonProblem
	expertiseConsoleInfoModel.IsDeviceHasUSBProblem = expertiseConsoleInfo.IsDeviceHasUSBProblem
	expertiseConsoleInfoModel.IsControllerHasUSBProblem = expertiseConsoleInfo.IsControllerHasUSBProblem
	expertiseConsoleInfoModel.IsDeviceHasHeatProblem = expertiseConsoleInfo.IsDeviceHasHeatProblem
	expertiseConsoleInfoModel.IsDeviceHasHighSoundProblem = expertiseConsoleInfo.IsDeviceHasHighSoundProblem
	expertiseConsoleInfoModel.IsDeviceHasNetworkConnectionProblem = expertiseConsoleInfo.IsDeviceHasNetworkConnectionProblem
	expertiseConsoleInfoModel.IsDeviceHasOpricalDriverProblem = expertiseConsoleInfo.IsDeviceHasOpricalDriverProblem
	expertiseConsoleInfoModel.IsDeviceHasDiskReadingProblem = expertiseConsoleInfo.IsDeviceHasDiskReadingProblem
	expertiseConsoleInfoModel.IsDeviceHasEthernetProblem = expertiseConsoleInfo.IsDeviceHasEthernetProblem
	expertiseConsoleInfoModel.IsDeviceHasHDMMIProblem = expertiseConsoleInfo.IsDeviceHasHDMMIProblem
	expertiseConsoleInfoModel.IsDeviceHaveUSBInputsProblem = expertiseConsoleInfo.IsDeviceHaveUSBInputsProblem
	expertiseConsoleInfoModel.IsDeviceHasBluetoothProblem = expertiseConsoleInfo.IsDeviceHasBluetoothProblem

	err := e.expertiseConsoleInfoService.Create(&expertiseConsoleInfoModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Error while creating expertise console info", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.HandleResponseModel(true, "Expertise console info created successfully", nil, expertiseConsoleInfoModel)
	ctx.JSON(http.StatusOK, response)
}
