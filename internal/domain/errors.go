package domain

import "errors"

var (
	ErrSubsNotFound = errors.New("subscription is not found")
	ErrSubNotUnique = errors.New("subscription with the same service name and user ID already exists")
)
