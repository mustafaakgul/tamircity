package tech_service

type BrandRequest struct {
	Name        string               `json:"name"`
	IsActive    bool                 `json:"is_active"`
	DeviceTypes []*DeviceTypeRequest `json:"device_types"`
}

type BrandResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
