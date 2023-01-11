package web

type ExpertiseWatchInfoRequest struct {
	ReservationId    uint    `json:"reservation_id"`
	Invoice          bool    `json:"invoice"`
	Box              bool    `json:"box"`
	GuaranteeTerm    int     `json:"guarantee_term"`
	PhoneColor       string  `json:"phone_color"`
	Microphone       bool    `json:"microphone"`
	ScreenSize       float32 `json:"screen_size"`
	Memory           int     `json:"memory"`
	ScreenType       string  `json:"screen_type"`
	OSType           string  `json:"os_type"`
	OSTypeCompatiple string  `json:"os_type_compatiple"`
	ScreenResolution string  `json:"screen_resolution"`
	CPUFrequency     int     `json:"cpu_frequency"`
	BatteryCapacity  int     `json:"battery_capacity"`
	Wifi             bool    `json:"wifi"`
	Speaker          bool    `json:"speaker"`
	SIMSupport       bool    `json:"sim_support"`
	DustProof        bool    `json:"dust_proof"`
	NFC              bool    `json:"nfc"`

	IsScreenHasBrokenProblem      bool `json:"is_screen_has_broken_problem"`
	IsScreenHasObscurationProblem bool `json:"is_screen_has_obscuration_problem"`
	IsTouchScreenHasProblem       bool `json:"is_touch_screen_has_problem"`
	IsScreenHasDeadPixelPixel     bool `json:"is_screen_has_dead_pixel_pixel"`
	IsDeviceHasCaseProblem        bool `json:"is_device_has_case_problem"`
	IsSpeakerHasProblem           bool `json:"is_speaker_has_problem"`
	IsPowerButtonHasProblem       bool `json:"is_power_button_has_problem"`
	IsBatterySocketHasProblem     bool `json:"is_battery_socket_has_problem"`
	IsDeviceHasFreezingProblem    bool `json:"is_device_has_freezing_problem"`
	IsBluetoothHasProblem         bool `json:"is_bluetooth_has_problem"`
	IsWiFiHasProblem              bool `json:"is_wifi_has_problem"`
	IsMicrophoneHasProblem        bool `json:"is_microphone_has_problem"`
	IsDeviceHasCellularProblem    bool `json:"is_device_has_cellular_problem"`
}
