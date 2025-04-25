package domain

import (
	"math"
	"strings"
	"time"
)

type Transaction struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
}

func NewTransaction(uId string, description string, date time.Time, value float64) (*Transaction, error) {
	if err := validateTransactionID(uId); err != nil {
		return nil, err
	}
	if err := validateTransactionDescription(description); err != nil {
		return nil, err
	}
	if err := validateTransactionDate(date); err != nil {
		return nil, err
	}
	if err := validateTransactionValue(value); err != nil {
		return nil, err
	}

	roundedValue := math.Round(value*100) / 100

	return &Transaction{
		ID:          uId,
		Description: description,
		Date:        date,
		Value:       roundedValue,
	}, nil
}

func validateTransactionID(id string) error {
	if id == "" {
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
