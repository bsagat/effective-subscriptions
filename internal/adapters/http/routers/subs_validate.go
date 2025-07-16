package routers

import (
	"errors"
	"fmt"
	"submanager/internal/domain"
	"unicode"
)

// Validates subscription data
func validateSubs(subs domain.Subscription) error {
	errFormat := "missing required field %s"
	if len(subs.UserID) == 0 {
		return fmt.Errorf(errFormat, "user_id")
	}

	if len(subs.ServiceName) == 0 {
		return fmt.Errorf(errFormat, "service_name")
	}

	if subs.StartDate.IsZero() {
		return fmt.Errorf(errFormat, "start_date")
	}

	if subs.Price <= 0 {
		return fmt.Errorf("price field must be more than 0")
	}

	if !IsValidUUID(subs.UserID) {
		return errors.New("user_ID is not UUID format")
	}
	return nil
}

// ValidateSubscriptionParams validates the parameters of a subscription request.
func validateSubsParams(serviceName, userID string) error {
	errFormat := "missing required query parameter %s"
	if len(userID) == 0 {
		return fmt.Errorf(errFormat, "user_id")
	}

	if !IsValidUUID(userID) {
		return errors.New("user_ID is not UUID format")
	}

	if len(serviceName) == 0 {
		return fmt.Errorf(errFormat, "service_name")
	}
	return nil
}

// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// 185925eb-2114-4c2a-bae7-6fdafa58d1d5

// x — hex (0-9, a-f, A-F), length 36, 8,13,18,23 positions are hyphens
func IsValidUUID(s string) bool {
	if len(s) != 36 {
		return false
	}

	for i, c := range s {
		switch i {
		case 8, 13, 18, 23:
			if c != '-' {
				return false
			}
		default:
			if !isHexChar(c) {
				return false
			}
		}
	}
	return true
}

func isHexChar(c rune) bool {
	return unicode.IsDigit(c) ||
		('a' <= c && c <= 'f') ||
		('A' <= c && c <= 'F')
}
