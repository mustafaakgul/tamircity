package tech_service

type FixTypeResponse struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
