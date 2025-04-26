package entities

import (
	"math"
	"strings"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
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
		return domain.ErrTransactionEmptyID
	}
	return nil
}

func validateTransactionDescription(description string) error {
	description = strings.TrimSpace(description)
	if description == "" {
		return domain.ErrTransactionEmptyDescription
	}
	if len(description) > 50 {
		return domain.ErrTransactionTooLongDescription
	}
	return nil
}

func validateTransactionDate(date time.Time) error {
	if date.IsZero() {
		return domain.ErrTransactionInvalidDate
	}
	return nil
}

func validateTransactionValue(value float64) error {
	if value <= 0 {
		return domain.ErrTransactionInvalidValue
	}
	return nil
}
