package handler

import (
	"net/http"
	"strconv"

	"github.com/anthophora/tamircity/pkg/models/db"
	"github.com/anthophora/tamircity/pkg/models/web"
	"github.com/anthophora/tamircity/pkg/service"
	"github.com/anthophora/tamircity/pkg/utils"
	"github.com/gin-gonic/gin"
)

type expertisePhoneInfoHandler struct {
	expertisePhoneInfoService service.ExpertisePhoneInfoService
}

type ExpertisePhoneInfoHandler interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

func NewExpertisePhoneInfoHandler(expertisePhoneInfoService service.ExpertisePhoneInfoService) ExpertisePhoneInfoHandler {
	return &expertisePhoneInfoHandler{
		expertisePhoneInfoService: expertisePhoneInfoService,
	}
}

func (e *expertisePhoneInfoHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := utils.HandleResponseModel(false, "Wrong Params", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	expertisePhoneInfo, err := e.expertisePhoneInfoService.GetByID(id)
	if err != nil {
		response := utils.HandleResponseModel(false, "Expertise Phone Info could not be found", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.HandleResponseModel(true, "Expertise Phone Info found successfully", nil, expertisePhoneInfo)
	ctx.JSON(http.StatusOK, response)
}

func (e *expertisePhoneInfoHandler) Create(ctx *gin.Context) {
	var expertisePhoneInfo web.ExpertisePhoneInfoRequest
	if err := ctx.ShouldBindJSON(&expertisePhoneInfo); err != nil {
		response := utils.HandleResponseModel(false, "Invalid JSON body", err, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var expertisePhoneInfoModel db.ExpertisePhoneInfo
	expertisePhoneInfoModel.ReservationId = expertisePhoneInfo.ReservationId
	expertisePhoneInfoModel.Invoice = expertisePhoneInfo.Invoice
	expertisePhoneInfoModel.Box = expertisePhoneInfo.Box
	expertisePhoneInfoModel.GuaranteeTerm = expertisePhoneInfo.GuaranteeTerm
	expertisePhoneInfoModel.Color = expertisePhoneInfo.Color
	expertisePhoneInfoModel.IMEIRegistration = expertisePhoneInfo.IMEIRegistration
	expertisePhoneInfoModel.EDevletRegistration = expertisePhoneInfo.EDevletRegistration
	expertisePhoneInfoModel.ScreenSize = expertisePhoneInfo.ScreenSize
	expertisePhoneInfoModel.ScreenTechnology = expertisePhoneInfo.ScreenTechnology
	expertisePhoneInfoModel.ScreenResolution = expertisePhoneInfo.ScreenResolution
	expertisePhoneInfoModel.ScratchResistance = expertisePhoneInfo.ScratchResistance
	expertisePhoneInfoModel.CPUModel = expertisePhoneInfo.CPUModel
	expertisePhoneInfoModel.CPUFrequency = expertisePhoneInfo.CPUFrequency
	expertisePhoneInfoModel.Ram = expertisePhoneInfo.Ram
	expertisePhoneInfoModel.OSType = expertisePhoneInfo.OSType
	expertisePhoneInfoModel.OSTypeVersion = expertisePhoneInfo.OSTypeVersion
	expertisePhoneInfoModel.CPUCoreNumber = expertisePhoneInfo.CPUCoreNumber
	expertisePhoneInfoModel.CameraResolution = expertisePhoneInfo.CameraResolution
	expertisePhoneInfoModel.FrontCameraResolution = expertisePhoneInfo.FrontCameraResolution
	expertisePhoneInfoModel.VideoRecordResolution = expertisePhoneInfo.VideoRecordResolution
	expertisePhoneInfoModel.VideoFPS = expertisePhoneInfo.VideoFPS
	expertisePhoneInfoModel.FaceRecognition = expertisePhoneInfo.FaceRecognition
	expertisePhoneInfoModel.SlowMotionVideo = expertisePhoneInfo.SlowMotionVideo
	expertisePhoneInfoModel.CameraAI = expertisePhoneInfo.CameraAI
	expertisePhoneInfoModel.Timer = expertisePhoneInfo.Timer
	expertisePhoneInfoModel.AutomaticFocus = expertisePhoneInfo.AutomaticFocus
	expertisePhoneInfoModel.GeographicLocation = expertisePhoneInfo.GeographicLocation
	expertisePhoneInfoModel.VoiceControl = expertisePhoneInfo.VoiceControl
	expertisePhoneInfoModel.InternalStorage = expertisePhoneInfo.InternalStorage
	expertisePhoneInfoModel.ExternalStorage = expertisePhoneInfo.ExternalStorage
	expertisePhoneInfoModel.MaxExternalStorage = expertisePhoneInfo.MaxExternalStorage
	expertisePhoneInfoModel.BatteryType = expertisePhoneInfo.BatteryType
	expertisePhoneInfoModel.BatteryCapacity = expertisePhoneInfo.BatteryCapacity
	expertisePhoneInfoModel.BatteryWirelessCharge = expertisePhoneInfo.BatteryWirelessCharge
	expertisePhoneInfoModel.BatteryFastCharge = expertisePhoneInfo.BatteryFastCharge
	expertisePhoneInfoModel.BatteryWirelessFastCharge = expertisePhoneInfo.BatteryWirelessFastCharge
	expertisePhoneInfoModel.BatteryDetachable = expertisePhoneInfo.BatteryDetachable
	expertisePhoneInfoModel.WiFiFrequency = expertisePhoneInfo.WiFiFrequency
	expertisePhoneInfoModel.NFC = expertisePhoneInfo.NFC
	expertisePhoneInfoModel.G5Support = expertisePhoneInfo.G5Support
	expertisePhoneInfoModel.ReleaseYear = expertisePhoneInfo.ReleaseYear
	expertisePhoneInfoModel.ResistanceofWater = expertisePhoneInfo.ResistanceofWater
	expertisePhoneInfoModel.ResistanceofDust = expertisePhoneInfo.ResistanceofDust
	expertisePhoneInfoModel.Fingerprint = expertisePhoneInfo.Fingerprint
	expertisePhoneInfoModel.DoubleSim = expertisePhoneInfo.DoubleSim
	expertisePhoneInfoModel.AnTuTuScore = expertisePhoneInfo.AnTuTuScore
	expertisePhoneInfoModel.IsScreenHasBrokenProblem = expertisePhoneInfo.IsScreenHasBrokenProblem
	expertisePhoneInfoModel.IsScreenHasObscurationProblem = expertisePhoneInfo.IsScreenHasObscurationProblem
	expertisePhoneInfoModel.IsTouchScreenHasProblem = expertisePhoneInfo.IsTouchScreenHasProblem
	expertisePhoneInfoModel.IsScreenHasDeadPixelPixel = expertisePhoneInfo.IsScreenHasDeadPixelProblem
	expertisePhoneInfoModel.IsDeviceHasCaseProblem = expertisePhoneInfo.IsDeviceHasCaseProblem
	expertisePhoneInfoModel.IsDeviceHasCoverProblem = expertisePhoneInfo.IsDeviceHasCoverProblem
	expertisePhoneInfoModel.IsDeviceHaveCamerasProblem = expertisePhoneInfo.IsDeviceHaveCamerasProblem
	expertisePhoneInfoModel.IsDeviceHaveSpeakerProblem = expertisePhoneInfo.IsDeviceHaveSpeakerProblem
	expertisePhoneInfoModel.IsDeviceHasHighHeatProblem = expertisePhoneInfo.IsDeviceHasHighHeatProblem
	expertisePhoneInfoModel.IsDeviceHasChargeSocketProblem = expertisePhoneInfo.IsDeviceHasChargeSocketProblem
	expertisePhoneInfoModel.IsDeviceHasPowerButtonProblem = expertisePhoneInfo.IsDeviceHasPowerButtonProblem
	expertisePhoneInfoModel.IsDeviceHasOpenedCaseProblem = expertisePhoneInfo.IsDeviceHasOpenedCaseProblem
	expertisePhoneInfoModel.IsDeviceHasSideButtonProblem = expertisePhoneInfo.IsDeviceHasSideButtonProblem
	expertisePhoneInfoModel.IsDeviceHasFreezingProblem = expertisePhoneInfo.IsDeviceHasFreezingProblem
	expertisePhoneInfoModel.IsDeviceHasBluetoothProblem = expertisePhoneInfo.IsDeviceHasBluetoothProblem
	expertisePhoneInfoModel.IsDeviceHasWiFiProblem = expertisePhoneInfo.IsDeviceHasWiFiProblem
	expertisePhoneInfoModel.IsDeviceHasMicrophoneProblem = expertisePhoneInfo.IsDeviceHasMicrophoneProblem
	expertisePhoneInfoModel.IsDeviceHasCellularProblem = expertisePhoneInfo.IsDeviceHasCellularProblem
	expertisePhoneInfoModel.IsDeviceHasSoundTransferProblem = expertisePhoneInfo.IsDeviceHasSoundTransferProblem

	err := e.expertisePhoneInfoService.Create(&expertisePhoneInfoModel)
	if err != nil {
		response := utils.HandleResponseModel(false, "Error while creating expertise phone info", err, nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := utils.HandleResponseModel(true, "Expertise phone info created successfully", nil, expertisePhoneInfoModel)
	ctx.JSON(http.StatusOK, response)
}
