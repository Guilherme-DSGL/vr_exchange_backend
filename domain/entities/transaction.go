package entities

import (
	"time"
)

type Transaction struct {
	ID          string    `json:"id" validate:"required"`
	Description string    `json:"description" validate:"required,max=50"`
	Date        time.Time `json:"date" validate:"required"`
	Value       float64   `json:"value" validate:"required,gt=0"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
