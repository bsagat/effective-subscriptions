package domain

import (
	"errors"
)

var (
	ErrSubsNotFound  = errors.New("subscription is not found")
	ErrSubNotUnique  = errors.New("subscription with the same service name and user ID already exists")
	ErrInvalidJSON   = errors.New("invalid JSON data")
	ErrInvalidUserID = errors.New("user_ID is not UUID format")
	ErrInvalidDate   = errors.New("start_date must be after end_date")
	ErrPriceField    = errors.New("price field must be more than 0")
)
