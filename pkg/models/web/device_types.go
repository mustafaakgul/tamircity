package web

type DeviceTypeRequest struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	IsActive         bool   `json:"is_active"`
}

type DeviceTypeResponse struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
}
