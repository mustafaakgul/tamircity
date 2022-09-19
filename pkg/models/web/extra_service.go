package web

type ExtraServiceResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
