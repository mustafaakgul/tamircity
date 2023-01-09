package db

import "gorm.io/gorm"

type ExpertiseWatchInfo struct {
	gorm.Model
	ReservationId    uint
	Reservation      *Reservation `gorm:"foreignkey:ReservationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invoice          bool         //Fatura var mi
	Box              bool         //Kutu
	GuaranteeTerm    int          //1-24
	PhoneColor       string
	Microphone       bool
	ScreenSize       float32
	Memory           int
	ScreenType       string
	OSType           string
	OSTypeCompatiple string
	ScreenResolution string
	CPUFrequency     int
	BatteryCapacity  int
	Wifi             bool
	Speaker          bool
	SIMSupport       bool
	DustProof        bool
	NFC              bool

	IsScreenHasBrokenProblem      bool
	IsScreenHasObscurationProblem bool
	IsTouchScreenHasProblem       bool
	IsScreenHasDeadPixelPixel     bool
	IsDeviceHasCaseProblem        bool
	IsSpeakerHasProblem           bool
	IsPowerButtonHasProblem       bool
	IsBatterySocketHasProblem     bool
	IsDeviceHasFreezingProblem    bool
	IsBluetoothHasProblem         bool
	IsWiFiHasProblem              bool
	IsMicrophoneHasProblem        bool
	IsDeviceHasCellularProblem    bool
}
