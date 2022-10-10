package db

import (
	"gorm.io/gorm"
)

type ExpertiseGameConsoleInfo struct {
	gorm.Model
	ReservationId    uint
	Reservation      *Reservation `gorm:"foreignkey:ReservationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invoice          bool         //Fatura var mi
	Box              bool         //Kutu
	GuaranteeTerm    int          //1-24
	Color            string
	Platform         string
	StorageCapacity  int
	OsType           string
	CPUModel         string
	FpsValue         int
	HDRSupport       bool
	Ram              int
	RamType          string
	GameResolution   string
	VideoResolution  string
	NetworkType      string
	Ethernet         bool
	EthernetSpeed    int
	HDMIStandart     string
	USBInput         int
	USBInputNumber   int
	USBVersion       string
	Bluetooth        bool
	BluetoothVersion string
	VoiceSupport     bool
	Controller       bool
	ControllerNumber int
	Game             bool
	GameNumber       int

	IsBoxHasProblem                     bool
	IsBoxHasLightProblem                bool
	IsControllerHasLightProblem         bool
	IsControllerHasVibrationProblem     bool
	IsControllerHasAnalogProglem        bool
	IsControllerHasButtonProblem        bool
	IsDeviceHasButtonProblem            bool
	IsDeviceHasUSBProblem               bool
	IsControllerHasUSBProblem           bool
	IsDeviceHasHeatProblem              bool
	IsDeviceHasHighSoundProblem         bool
	IsDeviceHasNetworkConnectionProblem bool
	IsDeviceHasOpricalDriverProblem     bool
	IsDeviceHasDiskReadingProblem       bool
	IsDeviceHasEthernetProblem          bool
	IsDeviceHasHDMMIProblem             bool
	IsDeviceHaveUSBInputsProblem        bool
	IsDeviceHasBluetoothProblem         bool
}
