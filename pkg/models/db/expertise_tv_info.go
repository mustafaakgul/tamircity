package db

import "gorm.io/gorm"

type ExpertiseTvInfo struct {
	gorm.Model
	ReservationId             uint
	Reservation               *Reservation `gorm:"foreignkey:ReservationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invoice                   bool         //Fatura var mi
	Box                       bool         //Kutu
	GuaranteeTerm             int          //1-24
	GuaranteeTermType         string
	Color                     string
	SmartTv                   bool
	PanelTechnology           string
	CurvedScreen              bool
	OSType                    string
	InternalSatelliteReceiver bool
	ScreenSize                float32
	Ambilight                 bool
	HDR                       bool
	FullArray                 bool
	ScreenResolution          string
	RefreshRate               int
	ReleaseYear               int
	ScreenShare               bool
	WiFi                      bool
	RFInput                   bool
	LAN                       bool
	HeadphoneJack             bool
	Bluetooth                 bool
	HDMI                      bool
	USB                       bool

	IsScreenHasObscurationProblem      bool
	IsScreenHasDeadPixelProblem        bool
	IsScreenHasBrokenProblem           bool
	IsScreenHasDownfallProblem         bool
	IsScreenHasBandProblem             bool
	IsScreenHasFreezingProblem         bool
	IsDeviceHasHighHeatProblem         bool
	IsDeviceHasSoundProblem            bool
	IsDeviceHasInternalReceiverProblem bool
	IsDeviceHasSpeakerProblem          bool
	IsScreenShareHasProblem            bool
	IsWiFiHasProblem                   bool
	IsRFInputHasProblem                bool
	IsEthernetHasProblem               bool
	IsHeadphoneJackHasProblem          bool
	IsBluetoothHasProblem              bool
	IsHDMIHasProblem                   bool
	IsUSBHasProblem                    bool
	IsRemoteControlHasProblem          bool
}
