package db

import "gorm.io/gorm"

type ExpertisePcInfo struct {
	gorm.Model
	ReservationId                 uint
	Reservation                   *Reservation `gorm:"foreignkey:ReservationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invoice                       bool         //Fatura var mi
	Box                           bool         //Kutu
	GuaranteeTerm                 int          //1-24
	PhoneColor                    string
	ScrenSize                     float32 // TODO: Refactor
	ScreenResolution              string  // TODO: Refactor
	ScreenResolutionType          string  // TODO: Refactor
	ScreenPanelType               string  // TODO: Refactor
	OperatingSystem               string  // TODO: Refactor
	CardReader                    bool
	Camera                        bool
	FingerPrintReader             bool
	TouchScreen                   bool
	KeyboardBacklight             bool
	CPU                           string
	CPUSerie                      string
	CPUBaseFrequency              float32
	CPUCoreNumber                 int
	CPURam                        int
	IntelTurboBoost               bool
	Ram                           int
	RamFrequency                  string
	RamType                       string
	HDDCapacity                   int
	SSD                           bool
	SSDCapacity                   int
	ExternalGraphicsCard          bool // Harici Ekran Karti
	ExternalGraphicsCardBrand     string
	ExternalGraphicsCardSeries    string
	ExternalGraphicsCardBit       int
	ExternalGraphicsCardCoreSpeed int
	// Dahili Ekran Karti
	InternalGraphicCardCPU     string // ????
	GraphicBaseFrequency       int
	GraphicMaxDynamicFrequency int
	FourK                      bool
	DirectX                    bool
	OpenGL                     float32
	SuppertedScreenNumber      int
	Ethernet                   bool
	EthernetVersion            bool
	WifiVersion                string
	HDMI                       bool
	Bluetooth                  bool
	USB2Number                 int
	USB3Number                 int
	USBTypeCNumber             int

	IsScreenHasBrokenProblem              bool
	IsScreenHasLossProblem                bool
	IsScreenHasDeadPixelPixel             bool
	IsDeviceHasCaseProblem                bool
	IsDeviceHasCoverProblem               bool
	IsKeyboardHasBrokenKeyProblem         bool
	IsKeyboardHasFunctionlessKeyProblem   bool
	IsDeviceHasChargeSocketProblem        bool
	IsDeviceHasTouchPadProblem            bool
	IsDeviceHasSpeakerProblem             bool
	IsDeviceHasHighHeatProblem            bool
	IsDeviceHasHighSoundProblem           bool
	IsDeviceHasDVDDriverProblem           bool
	IsDeviceHasUSBInputProblem            bool
	IsDeviceHasUSBTypeCProblem            bool
	IsDeviceHasCardReaderProblem          bool
	IsDeviceHasCameraProblem              bool
	IsDeviceHasFingerPrintProblem         bool
	IsDeviceHasMotherboardProblem         bool
	IsDeviceHasRAMProblem                 bool
	IsDeviceHasHDDProblem                 bool
	IsDeviceHasSSDProblem                 bool
	IsDeviceHasExternalGraphicCardProblem bool
	IsDeviceHasInternalGraphicCardProblem bool
	IsDeviceHasOpticalReaderProblem       bool
	IsDeviceHasEthernetConnectionProblem  bool
	IsDeviceHasWiFiProblem                bool
	IsDeviceHasHDMIProblem                bool
	IsDeviceHasBluetoothProblem           bool
	IsDeviceHasHeadPhoneSocketProblem     bool
	IsDeviceHasTouchScreenProblem         bool
	IsDeviceHasKeyboardBacklightProblem   bool
}

// TODO: Refactor
type PhoneColor int

const (
	AcikGri    PhoneColor = 0
	Altin                 = 1
	Beyaz                 = 2
	Bordo                 = 3
	Bronz                 = 4
	Gri                   = 5
	Gumus                 = 6
	Kahverengi            = 7
	Kirmizi               = 8
	Koyugri               = 9
	KoyuMavi              = 10
	Lacivert              = 11
	Mavi                  = 12
	Mor                   = 13
	Pembe                 = 14
	Siyah                 = 15
	turuncu               = 16
	Yesil                 = 17
)
