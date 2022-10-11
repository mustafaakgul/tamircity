package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type expertiseTvInfoHandler struct {
	expertiseTvInfoService service.ExpertiseTvInfoService
}

type ExpertiseTvInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertiseTvInfoHandler(expertiseTvInfoService service.ExpertiseTvInfoService) ExpertiseTvInfoHandler {
	return &expertiseTvInfoHandler{
		expertiseTvInfoService: expertiseTvInfoService,
	}
}

func (e *expertiseTvInfoHandler) Create(ctx *gin.Context) {
	var expertiseTvInfo web.ExpertiseTvInfoRequest
	if err := ctx.ShouldBindJSON(&expertiseTvInfo); err != nil {
		response := utils.HandleResponseModel(false, "Invalid request body", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertiseTvInfoModel db.ExpertiseTvInfo
	expertiseTvInfoModel.ReservationId = expertiseTvInfo.ReservationId
	expertiseTvInfoModel.Invoice = expertiseTvInfo.Invoice
	expertiseTvInfoModel.Box = expertiseTvInfo.Box
	expertiseTvInfoModel.GuaranteeTerm = expertiseTvInfo.GuaranteeTerm
	expertiseTvInfoModel.GuaranteeTermType = expertiseTvInfo.GuaranteeTermType
	expertiseTvInfoModel.Color = expertiseTvInfo.Color
	expertiseTvInfoModel.SmartTv = expertiseTvInfo.SmartTv
	expertiseTvInfoModel.PanelTechnology = expertiseTvInfo.PanelTechnology
	expertiseTvInfoModel.CurvedScreen = expertiseTvInfo.CurvedScreen
	expertiseTvInfoModel.OSType = expertiseTvInfo.OSType
	expertiseTvInfoModel.InternalSatelliteReceiver = expertiseTvInfo.InternalSatelliteReceiver
	expertiseTvInfoModel.ScreenSize = expertiseTvInfo.ScreenSize
	expertiseTvInfoModel.Ambilight = expertiseTvInfo.Ambilight
	expertiseTvInfoModel.HDR = expertiseTvInfo.HDR
	expertiseTvInfoModel.FullArray = expertiseTvInfo.FullArray
	expertiseTvInfoModel.ScreenResolution = expertiseTvInfo.ScreenResolution
	expertiseTvInfoModel.RefreshRate = expertiseTvInfo.RefreshRate
	expertiseTvInfoModel.ReleaseYear = expertiseTvInfo.ReleaseYear
	expertiseTvInfoModel.ScreenShare = expertiseTvInfo.ScreenShare
	expertiseTvInfoModel.WiFi = expertiseTvInfo.WiFi
	expertiseTvInfoModel.RFInput = expertiseTvInfo.RFInput
	expertiseTvInfoModel.LAN = expertiseTvInfo.LAN
	expertiseTvInfoModel.HeadphoneJack = expertiseTvInfo.HeadphoneJack
	expertiseTvInfoModel.Bluetooth = expertiseTvInfo.Bluetooth
	expertiseTvInfoModel.HDMI = expertiseTvInfo.HDMI
	expertiseTvInfoModel.USB = expertiseTvInfo.USB
	expertiseTvInfoModel.IsScreenHasObscurationProblem = expertiseTvInfo.IsScreenHasObscurationProblem
	expertiseTvInfoModel.IsScreenHasDeadPixelProblem = expertiseTvInfo.IsScreenHasDeadPixelProblem
	expertiseTvInfoModel.IsScreenHasBrokenProblem = expertiseTvInfo.IsScreenHasBrokenProblem
	expertiseTvInfoModel.IsScreenHasDownfallProblem = expertiseTvInfo.IsScreenHasDownfallProblem
	expertiseTvInfoModel.IsScreenHasBandProblem = expertiseTvInfo.IsScreenHasBandProblem
	expertiseTvInfoModel.IsScreenHasFreezingProblem = expertiseTvInfo.IsScreenHasFreezingProblem
	expertiseTvInfoModel.IsDeviceHasHighHeatProblem = expertiseTvInfo.IsDeviceHasHighHeatProblem
	expertiseTvInfoModel.IsDeviceHasSoundProblem = expertiseTvInfo.IsDeviceHasSoundProblem
	expertiseTvInfoModel.IsDeviceHasInternalReceiverProblem = expertiseTvInfo.IsDeviceHasInternalReceiverProblem
	expertiseTvInfoModel.IsDeviceHasSpeakerProblem = expertiseTvInfo.IsDeviceHasSpeakerProblem
	expertiseTvInfoModel.IsScreenShareHasProblem = expertiseTvInfo.IsScreenShareHasProblem
	expertiseTvInfoModel.IsWiFiHasProblem = expertiseTvInfo.IsWiFiHasProblem
	expertiseTvInfoModel.IsRFInputHasProblem = expertiseTvInfo.IsRFInputHasProblem
	expertiseTvInfoModel.IsEthernetHasProblem = expertiseTvInfo.IsEthernetHasProblem
	expertiseTvInfoModel.IsHeadphoneJackHasProblem = expertiseTvInfo.IsHeadphoneJackHasProblem
	expertiseTvInfoModel.IsBluetoothHasProblem = expertiseTvInfo.IsBluetoothHasProblem
	expertiseTvInfoModel.IsHDMIHasProblem = expertiseTvInfo.IsHDMIHasProblem
	expertiseTvInfoModel.IsUSBHasProblem = expertiseTvInfo.IsUSBHasProblem
	expertiseTvInfoModel.IsRemoteControlHasProblem = expertiseTvInfo.IsRemoteControlHasProblem

	err := e.expertiseTvInfoService.Create(&expertiseTvInfoModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Error while creating expertise tv info", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.HandleResponseModel(true, "Expertise Tv info created successfully", nil, expertiseTvInfoModel)
	ctx.JSON(http.StatusOK, response)
}
