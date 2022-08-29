package web

type BrandRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type BrandResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
