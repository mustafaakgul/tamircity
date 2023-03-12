package web

import (
	"github.com/anthophora/tamircity/pkg/models/db"
	"time"
)

type PaymentInfoRequest struct {
	PaymentDate   time.Time      `json:"payment_date"`
	Amount        int            `json:"amount"`
	PaymentType   db.PaymentType `json:"payment_type"`
	ReservationID int            `json:"reservation_id"`
}

type PaymentInfoResponse struct {
	Id          uint           `json:"id"`
	PaymentDate time.Time      `json:"payment_date"`
	Amount      int            `json:"amount"`
	PaymentType db.PaymentType `json:"payment_type"`
}
