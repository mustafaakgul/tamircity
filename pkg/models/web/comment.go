package web

import "time"

type CommentRequest struct {
	ExpertiseServiceID int       `json:"expertise_service_id"`
	CommentOwner       string    `json:"comment_owner"`
	Comment            string    `json:"comment"`
	CommentDate        time.Time `json:"comment_date"`
	Rate               int       `json:"rate"`
}
