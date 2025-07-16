package domain

import "time"

type Subscription struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
