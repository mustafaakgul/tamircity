package web

type ExpertiseConsoleInfoRequest struct {
	ReservationId    uint   `json:"reservation_id"`
	Invoice          bool   `json:"invoice"`
	Box              bool   `json:"box"`
	GuaranteeTerm    int    `json:"guarantee_term"`
	Color            string `json:"color"`
	Platform         string `json:"platform"`
	StorageCapacity  int    `json:"storage_capacity"`
	OsType           string `json:"os_type"`
	CPUModel         string `json:"cpu_model"`
	FpsValue         int    `json:"fps_value"`
	HDRSupport       bool   `json:"hdr_support"`
	Ram              int    `json:"ram"`
	RamType          string `json:"ram_type"`
	GameResolution   string `json:"game_resolution"`
	VideoResolution  string `json:"video_resolution"`
	NetworkType      string `json:"network_type"`
	Ethernet         bool   `json:"ethernet"`
	EthernetSpeed    int    `json:"ethernet_speed"`
	HDMIStandart     string `json:"hdmi_standart"`
	USBInput         int    `json:"usb_input"`
	USBInputNumber   int    `json:"usb_input_number"`
	USBVersion       string `json:"usb_version"`
	Bluetooth        bool   `json:"bluetooth"`
	BluetoothVersion string `json:"bluetooth_version"`
	VoiceSupport     bool   `json:"voice_support"`
	Controller       bool   `json:"controller"`
	ControllerNumber int    `json:"controller_number"`
	Game             bool   `json:"game"`
	GameNumber       int    `json:"game_number"`

	IsBoxHasProblem                     bool `json:"is_box_has_has_problem"`
	IsBoxHasLightProblem                bool `json:"is_box_has_light_problem"`
	IsControllerHasLightProblem         bool `json:"is_controller_has_light_problem"`
	IsControllerHasVibrationProblem     bool `json:"is_controller_has_vibration_problem"`
	IsControllerHasAnalogProglem        bool `json:"is_controller_has_analog_problem"`
	IsControllerHasButtonProblem        bool `json:"is_controller_has_button_problem"`
	IsDeviceHasButtonProblem            bool `json:"is_device_has_button_problem"`
	IsDeviceHasUSBProblem               bool `json:"is_device_has_usb_problem"`
	IsControllerHasUSBProblem           bool `json:"is_controller_has_usb_problem"`
	IsDeviceHasHeatProblem              bool `json:"is_device_has_heat_problem"`
	IsDeviceHasHighSoundProblem         bool `json:"is_device_has_hight_sound_problem"`
	IsDeviceHasNetworkConnectionProblem bool `json:"is_device_has_network_connection_problem"`
	IsDeviceHasOpricalDriverProblem     bool `json:"is_device_has_oprical_driver_problem"`
	IsDeviceHasDiskReadingProblem       bool `json:"is_device_has_disk_reading_problem"`
	IsDeviceHasEthernetProblem          bool `json:"is_device_has_ethernet_problem"`
	IsDeviceHasHDMMIProblem             bool `json:"is_device_has_hdmi_problem"`
	IsDeviceHaveUSBInputsProblem        bool `json:"is_device_has_usb_inputs_problem"`
	IsDeviceHasBluetoothProblem         bool `json:"is_device_has_bluetooth_problem"`
}
