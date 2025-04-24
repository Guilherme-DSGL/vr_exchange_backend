package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func validateTransactionID(id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrTransactionEmptyID
	}
	return nil
}

func validateTransactionDescription(description string) error {
	description = strings.TrimSpace(description)
	if description == "" {
		return ErrTransactionEmptyDescription
	}
	if len(description) > 50 {
		return ErrTransactionTooLongDescription
	}
	return nil
}

func validateTransactionDate(date time.Time) error {
	if date.IsZero() {
		return ErrTransactionInvalidDate
	}
	return nil
}

func validateTransactionValue(value float64) error {
	if value <= 0 {
		return ErrTransactionInvalidValue
	}
	return nil
}
