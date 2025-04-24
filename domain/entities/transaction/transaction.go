package domain

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
}

func NewTransaction(id uuid.UUID, description string, date time.Time, value float64) (*Transaction, error) {
	if err := validateTransactionID(id); err != nil {
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
		ID:          id,
		Description: description,
		Date:        date,
		Value:       roundedValue,
	}, nil
}
