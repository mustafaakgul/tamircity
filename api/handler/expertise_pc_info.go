package handler

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type expertisePcInfoHandler struct {
	expertisePcInfoService service.ExpertisePcInfoService
}

type ExpertisePcInfoHandler interface {
	Create(ctx *gin.Context)
}

func NewExpertisePcInfoHandler(expertisePcInfoService service.ExpertisePcInfoService) ExpertisePcInfoHandler {
	return &expertisePcInfoHandler{
		expertisePcInfoService: expertisePcInfoService,
	}
}

func (e *expertisePcInfoHandler) Create(ctx *gin.Context) {
	var expertisePcInfo web.ExpertisePcInfoRequest
	if err := ctx.ShouldBindJSON(&expertisePcInfo); err != nil {
		response := utils.HandleResponseModel(false, "Invalid JSON body", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertisePcInfoModel db.ExpertisePcInfo
	expertisePcInfoModel.ReservationId = expertisePcInfo.ReservationId
	expertisePcInfoModel.Invoice = expertisePcInfo.Invoice
	expertisePcInfoModel.Box = expertisePcInfo.Box
	expertisePcInfoModel.GuaranteeTerm = expertisePcInfo.GuaranteeTerm
	expertisePcInfoModel.PhoneColor = expertisePcInfo.PhoneColor
	expertisePcInfoModel.ScrenSize = expertisePcInfo.ScrenSize
	expertisePcInfoModel.ScreenResolution = expertisePcInfo.ScreenResolution
	expertisePcInfoModel.ScreenResolutionType = expertisePcInfo.ScreenResolutionType
	expertisePcInfoModel.ScreenPanelType = expertisePcInfo.ScreenPanelType
	expertisePcInfoModel.OperatingSystem = expertisePcInfo.OperatingSystem
	expertisePcInfoModel.CardReader = expertisePcInfo.CardReader
	expertisePcInfoModel.Camera = expertisePcInfo.Camera
	expertisePcInfoModel.FingerPrintReader = expertisePcInfo.FingerPrintReader
	expertisePcInfoModel.TouchScreen = expertisePcInfo.TouchScreen
	expertisePcInfoModel.KeyboardBacklight = expertisePcInfo.KeyboardBacklight
	expertisePcInfoModel.CPU = expertisePcInfo.CPU
	expertisePcInfoModel.CPUSerie = expertisePcInfo.CPUSerie
	expertisePcInfoModel.CPUBaseFrequency = expertisePcInfo.CPUBaseFrequency
	expertisePcInfoModel.CPUCoreNumber = expertisePcInfo.CPUCoreNumber
	expertisePcInfoModel.CPURam = expertisePcInfo.CPURam
	expertisePcInfoModel.IntelTurboBoost = expertisePcInfo.IntelTurboBoost
	expertisePcInfoModel.Ram = expertisePcInfo.Ram
	expertisePcInfoModel.RamFrequency = expertisePcInfo.RamFrequency
	expertisePcInfoModel.RamType = expertisePcInfo.RamType
	expertisePcInfoModel.HDDCapacity = expertisePcInfo.HDDCapacity
	expertisePcInfoModel.SSD = expertisePcInfo.SSD
	expertisePcInfoModel.SSDCapacity = expertisePcInfo.SSDCapacity
	expertisePcInfoModel.ExternalGraphicsCard = expertisePcInfo.ExternalGraphicsCard
	expertisePcInfoModel.ExternalGraphicsCardBrand = expertisePcInfo.ExternalGraphicsCardBrand
	expertisePcInfoModel.ExternalGraphicsCardSeries = expertisePcInfo.ExternalGraphicsCardSeries
	expertisePcInfoModel.ExternalGraphicsCardBit = expertisePcInfo.ExternalGraphicsCardBit
	expertisePcInfoModel.ExternalGraphicsCardCoreSpeed = expertisePcInfo.ExternalGraphicsCardCoreSpeed
	expertisePcInfoModel.InternalGraphicCardCPU = expertisePcInfo.InternalGraphicCardCPU
	expertisePcInfoModel.GraphicBaseFrequency = expertisePcInfo.GraphicBaseFrequency
	expertisePcInfoModel.GraphicMaxDynamicFrequency = expertisePcInfo.GraphicMaxDynamicFrequency
	expertisePcInfoModel.FourK = expertisePcInfo.FourK
	expertisePcInfoModel.DirectX = expertisePcInfo.DirectX
	expertisePcInfoModel.OpenGL = expertisePcInfo.OpenGL
	expertisePcInfoModel.SuppertedScreenNumber = expertisePcInfo.SuppertedScreenNumber
	expertisePcInfoModel.Ethernet = expertisePcInfo.Ethernet
	expertisePcInfoModel.EthernetVersion = expertisePcInfo.EthernetVersion
	expertisePcInfoModel.WifiVersion = expertisePcInfo.WifiVersion
	expertisePcInfoModel.HDMI = expertisePcInfo.HDMI
	expertisePcInfoModel.Bluetooth = expertisePcInfo.Bluetooth
	expertisePcInfoModel.USB2Number = expertisePcInfo.USB2Number
	expertisePcInfoModel.USB3Number = expertisePcInfo.USB3Number
	expertisePcInfoModel.USBTypeCNumber = expertisePcInfo.USBTypeCNumber
	expertisePcInfoModel.IsScreenHasBrokenProblem = expertisePcInfo.IsScreenHasBrokenProblem
	expertisePcInfoModel.IsScreenHasLossProblem = expertisePcInfo.IsScreenHasLossProblem
	expertisePcInfoModel.IsScreenHasDeadPixelPixel = expertisePcInfo.IsScreenHasDeadPixelPixel
	expertisePcInfoModel.IsDeviceHasCaseProblem = expertisePcInfo.IsDeviceHasCaseProblem
	expertisePcInfoModel.IsDeviceHasCoverProblem = expertisePcInfo.IsDeviceHasCoverProblem
	expertisePcInfoModel.IsKeyboardHasBrokenKeyProblem = expertisePcInfo.IsKeyboardHasBrokenKeyProblem
	expertisePcInfoModel.IsKeyboardHasFunctionlessKeyProblem = expertisePcInfo.IsKeyboardHasFunctionlessKeyProblem
	expertisePcInfoModel.IsDeviceHasChargeSocketProblem = expertisePcInfo.IsDeviceHasChargeSocketProblem
	expertisePcInfoModel.IsDeviceHasTouchPadProblem = expertisePcInfo.IsDeviceHasTouchPadProblem
	expertisePcInfoModel.IsDeviceHasSpeakerProblem = expertisePcInfo.IsDeviceHasSpeakerProblem
	expertisePcInfoModel.IsDeviceHasHighHeatProblem = expertisePcInfo.IsDeviceHasHighHeatProblem
	expertisePcInfoModel.IsDeviceHasHighSoundProblem = expertisePcInfo.IsDeviceHasHighSoundProblem
	expertisePcInfoModel.IsDeviceHasDVDDriverProblem = expertisePcInfo.IsDeviceHasDVDDriverProblem
	expertisePcInfoModel.IsDeviceHasUSBInputProblem = expertisePcInfo.IsDeviceHasUSBInputProblem
	expertisePcInfoModel.IsDeviceHasUSBTypeCProblem = expertisePcInfo.IsDeviceHasUSBTypeCProblem
	expertisePcInfoModel.IsDeviceHasCardReaderProblem = expertisePcInfo.IsDeviceHasCardReaderProblem
	expertisePcInfoModel.IsDeviceHasCameraProblem = expertisePcInfo.IsDeviceHasCameraProblem
	expertisePcInfoModel.IsDeviceHasFingerPrintProblem = expertisePcInfo.IsDeviceHasFingerPrintProblem
	expertisePcInfoModel.IsDeviceHasMotherboardProblem = expertisePcInfo.IsDeviceHasMotherboardProblem
	expertisePcInfoModel.IsDeviceHasRAMProblem = expertisePcInfo.IsDeviceHasRAMProblem
	expertisePcInfoModel.IsDeviceHasHDDProblem = expertisePcInfo.IsDeviceHasHDDProblem
	expertisePcInfoModel.IsDeviceHasSSDProblem = expertisePcInfo.IsDeviceHasSSDProblem
	expertisePcInfoModel.IsDeviceHasExternalGraphicCardProblem = expertisePcInfo.IsDeviceHasExternalGraphicCardProblem
	expertisePcInfoModel.IsDeviceHasInternalGraphicCardProblem = expertisePcInfo.IsDeviceHasInternalGraphicCardProblem
	expertisePcInfoModel.IsDeviceHasOpticalReaderProblem = expertisePcInfo.IsDeviceHasOpticalReaderProblem
	expertisePcInfoModel.IsDeviceHasEthernetConnectionProblem = expertisePcInfo.IsDeviceHasEthernetConnectionProblem
	expertisePcInfoModel.IsDeviceHasWiFiProblem = expertisePcInfo.IsDeviceHasWiFiProblem
	expertisePcInfoModel.IsDeviceHasHDMIProblem = expertisePcInfo.IsDeviceHasHDMIProblem
	expertisePcInfoModel.IsDeviceHasBluetoothProblem = expertisePcInfo.IsDeviceHasBluetoothProblem
	expertisePcInfoModel.IsDeviceHasHeadPhoneSocketProblem = expertisePcInfo.IsDeviceHasHeadPhoneSocketProblem
	expertisePcInfoModel.IsDeviceHasTouchScreenProblem = expertisePcInfo.IsDeviceHasTouchScreenProblem
	expertisePcInfoModel.IsDeviceHasKeyboardBacklightProblem = expertisePcInfo.IsDeviceHasKeyboardBacklightProblem

	err := e.expertisePcInfoService.Create(&expertisePcInfoModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Error while creating expertise pc info", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := utils.HandleResponseModel(true, "Expertise pc info created successfully", nil, expertisePcInfoModel)
	ctx.JSON(http.StatusOK, response)
}
