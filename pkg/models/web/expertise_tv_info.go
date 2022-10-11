package web

type ExpertiseTvInfoRequest struct {
	ReservationId             uint    `json:"reservation_id"`
	Invoice                   bool    `json:"invoice"`
	Box                       bool    `json:"box"`
	GuaranteeTerm             int     `json:"guarantee_term"`
	GuaranteeTermType         string  `json:"guarantee_term_type"`
	Color                     string  `json:"color"`
	SmartTv                   bool    `json:"smart_tv"`
	PanelTechnology           string  `json:"panel_technology"`
	CurvedScreen              bool    `json:"curved_screen"`
	OSType                    string  `json:"os_type"`
	InternalSatelliteReceiver bool    `json:"internal_satallite_recevier"`
	ScreenSize                float32 `json:"screen_size"`
	Ambilight                 bool    `json:"ambilight"`
	HDR                       bool    `json:"hdr"`
	FullArray                 bool    `json:"full_array"`
	ScreenResolution          string  `json:"screen_resolution"`
	RefreshRate               int     `json:"refresh_rate"`
	ReleaseYear               int     `json:"release_year"`
	ScreenShare               bool    `json:"screen_share"`
	WiFi                      bool    `json:"wifi"`
	RFInput                   bool    `json:"rfi_input"`
	LAN                       bool    `json:"lan"`
	HeadphoneJack             bool    `json:"head_phone_jack"`
	Bluetooth                 bool    `json:"bluetooth"`
	HDMI                      bool    `json:"hdmi"`
	USB                       bool    `json:"usb"`

	IsScreenHasObscurationProblem      bool `json:"is_screen_has_obscuration_problem"`
	IsScreenHasDeadPixelProblem        bool `json:"is_screen_has_dead_pixel_problem"`
	IsScreenHasBrokenProblem           bool `json:"is_screen_has_broken_problem"`
	IsScreenHasDownfallProblem         bool `json:"is_screen_has_down_fall_problem"`
	IsScreenHasBandProblem             bool `json:"is_screen_has_band_problem"`
	IsScreenHasFreezingProblem         bool `json:"is_screen_has_freezing_problem"`
	IsDeviceHasHighHeatProblem         bool `json:"is_device_has_high_heat_problem"`
	IsDeviceHasSoundProblem            bool `json:"is_device_has_sound_problem"`
	IsDeviceHasInternalReceiverProblem bool `json:"is_device_has_internal_receiver_problem"`
	IsDeviceHasSpeakerProblem          bool `json:"is_device_has_speaker_problem"`
	IsScreenShareHasProblem            bool `json:"is_screen_share_has_problem"`
	IsWiFiHasProblem                   bool `json:"is_wifi_has_problem"`
	IsRFInputHasProblem                bool `json:"is_rfi_input_has_problem"`
	IsEthernetHasProblem               bool `json:"is_ethernet_has_problem"`
	IsHeadphoneJackHasProblem          bool `json:"is_head_phone_jack_has_problem"`
	IsBluetoothHasProblem              bool `json:"is_bluetooth_has_problem"`
	IsHDMIHasProblem                   bool `json:"is_hdmi_has_problem"`
	IsUSBHasProblem                    bool `json:"is_usb_has_problem"`
	IsRemoteControlHasProblem          bool `json:"is_remote_control_has_problem"`
}
