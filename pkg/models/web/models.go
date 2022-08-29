package web

type ModelRequest struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	IsActive         bool   `json:"is_active"`
}

type ModelResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
