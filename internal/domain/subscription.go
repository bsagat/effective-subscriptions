package domain

import (
	"time"
)

type Subscription struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Summary struct {
	TotalPrice    int            `json:"total_price"`
	SubsCount     int            `json:"total_subscriptions"`
	PageNumber    int            `json:"page_number"`
	PageSize      int            `json:"page_size"`
	Subscriptions []Subscription `json:"subscriptions"`
}
