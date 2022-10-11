package web

type ExpertisePcInfoRequest struct {
	ReservationId                 uint    `json:"reservation_id"`
	Invoice                       bool    `json:"invoice"`
	Box                           bool    `json:"box"`
	GuaranteeTerm                 int     `json:"guarantee_term"`
	PhoneColor                    string  `json:"phone_color"`
	ScreenSize                    float32 `json:"screen_size"`
	ScreenResolution              string  `json:"screen_resolution"`
	ScreenResolutionType          string  `json:"screen_resolution_type"`
	ScreenPanelType               string  `json:"screen_panel_type"`
	OperatingSystem               string  `json:"operating_system"`
	CardReader                    bool    `json:"card_reader"`
	Camera                        bool    `json:"camera"`
	FingerPrintReader             bool    `json:"finger_print_reader"`
	TouchScreen                   bool    `json:"touch_screen"`
	KeyboardBacklight             bool    `json:"keyboard_backlight"`
	CPU                           string  `json:"cpu"`
	CPUSerie                      string  `json:"cpu_serie"`
	CPUBaseFrequency              float32 `json:"cpu_base_frequency"`
	CPUCoreNumber                 int     `json:"cpu_core_number"`
	CPURam                        int     `json:"cpu_ram"`
	IntelTurboBoost               bool    `json:"intel_turba_boost"`
	Ram                           int     `json:"ram"`
	RamFrequency                  string  `json:"ram_frequency"`
	RamType                       string  `json:"ram_type"`
	HDDCapacity                   int     `json:"hdd_capacity"`
	SSD                           bool    `json:"ssd"`
	SSDCapacity                   int     `json:"ssd_capacity"`
	ExternalGraphicsCard          bool    `json:"external_graphics_card"`
	ExternalGraphicsCardBrand     string  `json:"external_graphics_card_brand"`
	ExternalGraphicsCardSeries    string  `json:"external_graphics_card_series"`
	ExternalGraphicsCardBit       int     `json:"external_graphics_card_bit"`
	ExternalGraphicsCardCoreSpeed int     `json:"external_graphics_card_core_speed"`
	InternalGraphicCardCPU        string  `json:"internal_graphics_card_cpu"`
	GraphicBaseFrequency          int     `json:"graphic_base_frequency"`
	GraphicMaxDynamicFrequency    int     `json:"graphic_max_dynamic_frequency"`
	FourK                         bool    `json:"four_k"`
	DirectX                       bool    `json:"directx"`
	OpenGL                        float32 `json:"open_gl"`
	SuppertedScreenNumber         int     `json:"supperted_screen_number"`
	Ethernet                      bool    `json:"ethernet"`
	EthernetVersion               bool    `json:"ethernet_version"`
	WifiVersion                   string  `json:"wifi_version"`
	HDMI                          bool    `json:"hdmi"`
	Bluetooth                     bool    `json:"bluetooth"`
	USB2Number                    int     `json:"usb_2_number"`
	USB3Number                    int     `json:"usb_3_number"`
	USBTypeCNumber                int     `json:"usb_type_c_number"`

	IsScreenHasBrokenProblem              bool `json:"is_screen_has_broken_problem"`
	IsScreenHasLossProblem                bool `json:"is_screen_has_loss_problem"`
	IsScreenHasDeadPixelProblem           bool `json:"is_screen_has_dead_pixel_problem"`
	IsDeviceHasCaseProblem                bool `json:"is_device_has_case_problem"`
	IsDeviceHasCoverProblem               bool `json:"is_device_has_cover_problem"`
	IsKeyboardHasBrokenKeyProblem         bool `json:"is_device_has_broken_key_problem"`
	IsKeyboardHasFunctionlessKeyProblem   bool `json:"is_device_has_function_less_key_problem"`
	IsDeviceHasChargeSocketProblem        bool `json:"is_device_has_charge_socket_problem"`
	IsDeviceHasTouchPadProblem            bool `json:"is_device_has_touch_pad_problem"`
	IsDeviceHasSpeakerProblem             bool `json:"is_device_has_speaker_problem"`
	IsDeviceHasHighHeatProblem            bool `json:"is_device_has_high_heat_problem"`
	IsDeviceHasHighSoundProblem           bool `json:"is_device_has_high_sound_problem"`
	IsDeviceHasDVDDriverProblem           bool `json:"is_device_has_dvd_driver_problem"`
	IsDeviceHasUSBInputProblem            bool `json:"is_device_has_usb_input_problem"`
	IsDeviceHasUSBTypeCProblem            bool `json:"is_device_has_usb_type_c_problem"`
	IsDeviceHasCardReaderProblem          bool `json:"is_device_has_card_reader_problem"`
	IsDeviceHasCameraProblem              bool `json:"is_device_has_camera_problem"`
	IsDeviceHasFingerPrintProblem         bool `json:"is_device_has_finger_print_problem"`
	IsDeviceHasMotherboardProblem         bool `json:"is_device_has_mother_board_problem"`
	IsDeviceHasRAMProblem                 bool `json:"is_device_has_ram_problem"`
	IsDeviceHasHDDProblem                 bool `json:"is_device_has_hdd_problem"`
	IsDeviceHasSSDProblem                 bool `json:"is_device_has_ssd_problem"`
	IsDeviceHasExternalGraphicCardProblem bool `json:"is_device_has_external_graphic_card_problem"`
	IsDeviceHasInternalGraphicCardProblem bool `json:"is_device_has_internal_graphic_card_problem"`
	IsDeviceHasOpticalReaderProblem       bool `json:"is_device_has_optical_reader_problem"`
	IsDeviceHasEthernetConnectionProblem  bool `json:"is_device_has_ethernet_connection_problem"`
	IsDeviceHasWiFiProblem                bool `json:"is_device_has_wifi_problem"`
	IsDeviceHasHDMIProblem                bool `json:"is_device_has_hdmi_problem"`
	IsDeviceHasBluetoothProblem           bool `json:"is_device_has_bluetooth_problem"`
	IsDeviceHasHeadPhoneSocketProblem     bool `json:"is_device_has_head_phone_socket_problem"`
	IsDeviceHasTouchScreenProblem         bool `json:"is_device_has_touch_screen_problem"`
	IsDeviceHasKeyboardBacklightProblem   bool `json:"is_device_has_keyboard_backlight_problem"`
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
