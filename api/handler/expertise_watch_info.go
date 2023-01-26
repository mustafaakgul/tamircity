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

type expertiseWatchInfoHandler struct {
	expertiseWatchInfoService service.ExpertiseWatchInfoService
}

type ExpertiseWatchInfoHandler interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

func NewExpertiseWatchInfoHandler(expertiseWatchInfoService service.ExpertiseWatchInfoService) ExpertiseWatchInfoHandler {
	return &expertiseWatchInfoHandler{
		expertiseWatchInfoService: expertiseWatchInfoService,
	}
}

func (e *expertiseWatchInfoHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := utils.HandleResponseModel(false, "Wrong Params", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	expertiseWatchInfo, err := e.expertiseWatchInfoService.GetByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Watch Info could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Watch Info found successfully", nil, expertiseWatchInfo)
	ctx.JSON(http.StatusOK, response)
}

func (e *expertiseWatchInfoHandler) Create(ctx *gin.Context) {
	var expertiseWatchInfo web.ExpertiseWatchInfoRequest
	if err := ctx.ShouldBindJSON(&expertiseWatchInfo); err != nil {
		response := utils.HandleResponseModel(false, "Invalid JSON body", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseWatchInfoModel db.ExpertiseWatchInfo
	expertiseWatchInfoModel.ReservationId = expertiseWatchInfo.ReservationId
	expertiseWatchInfoModel.Invoice = expertiseWatchInfo.Invoice
	expertiseWatchInfoModel.Box = expertiseWatchInfo.Box
	expertiseWatchInfoModel.GuaranteeTerm = expertiseWatchInfo.GuaranteeTerm
	expertiseWatchInfoModel.PhoneColor = expertiseWatchInfo.PhoneColor
	expertiseWatchInfoModel.Microphone = expertiseWatchInfo.Microphone
	expertiseWatchInfoModel.ScreenSize = expertiseWatchInfo.ScreenSize
	expertiseWatchInfoModel.Memory = expertiseWatchInfo.Memory
	expertiseWatchInfoModel.ScreenType = expertiseWatchInfo.ScreenType
	expertiseWatchInfoModel.OSType = expertiseWatchInfo.OSType
	expertiseWatchInfoModel.OSTypeCompatiple = expertiseWatchInfo.OSTypeCompatiple
	expertiseWatchInfoModel.ScreenResolution = expertiseWatchInfo.ScreenResolution
	expertiseWatchInfoModel.CPUFrequency = expertiseWatchInfo.CPUFrequency
	expertiseWatchInfoModel.BatteryCapacity = expertiseWatchInfo.BatteryCapacity
	expertiseWatchInfoModel.Wifi = expertiseWatchInfo.Wifi
	expertiseWatchInfoModel.Speaker = expertiseWatchInfo.Speaker
	expertiseWatchInfoModel.SIMSupport = expertiseWatchInfo.SIMSupport
	expertiseWatchInfoModel.DustProof = expertiseWatchInfo.DustProof
	expertiseWatchInfoModel.NFC = expertiseWatchInfo.NFC

	expertiseWatchInfoModel.IsScreenHasBrokenProblem = expertiseWatchInfo.IsScreenHasBrokenProblem
	expertiseWatchInfoModel.IsScreenHasObscurationProblem = expertiseWatchInfo.IsScreenHasObscurationProblem
	expertiseWatchInfoModel.IsTouchScreenHasProblem = expertiseWatchInfo.IsTouchScreenHasProblem
	expertiseWatchInfoModel.IsScreenHasDeadPixelPixel = expertiseWatchInfo.IsScreenHasDeadPixelPixel
	expertiseWatchInfoModel.IsDeviceHasCaseProblem = expertiseWatchInfo.IsDeviceHasCaseProblem
	expertiseWatchInfoModel.IsSpeakerHasProblem = expertiseWatchInfo.IsSpeakerHasProblem
	expertiseWatchInfoModel.IsPowerButtonHasProblem = expertiseWatchInfo.IsPowerButtonHasProblem
	expertiseWatchInfoModel.IsBatterySocketHasProblem = expertiseWatchInfo.IsBatterySocketHasProblem
	expertiseWatchInfoModel.IsDeviceHasFreezingProblem = expertiseWatchInfo.IsDeviceHasFreezingProblem
	expertiseWatchInfoModel.IsBluetoothHasProblem = expertiseWatchInfo.IsBluetoothHasProblem
	expertiseWatchInfoModel.IsWiFiHasProblem = expertiseWatchInfo.IsWiFiHasProblem
	expertiseWatchInfoModel.IsMicrophoneHasProblem = expertiseWatchInfo.IsMicrophoneHasProblem
	expertiseWatchInfoModel.IsDeviceHasCellularProblem = expertiseWatchInfo.IsDeviceHasCellularProblem

	err := e.expertiseWatchInfoService.Create(&expertiseWatchInfoModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Error while creating expertise watch info", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.HandleResponseModel(true, "Expertise watch info created successfully", nil, expertiseWatchInfoModel)
	ctx.JSON(http.StatusOK, response)
}
