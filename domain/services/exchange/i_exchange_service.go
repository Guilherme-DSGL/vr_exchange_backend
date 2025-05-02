package services

import (
	"context"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
)

// Represents the business logic for handling exchange operations.
//
//go:generate mockery --name IExchangeService
type IExchangeService interface {
	// GetExchange retrieves an exchange transaction based on the provided exchange parameters.
	// Returns the exchange transaction and an error if the operation fails.
	// If not exist exchange in the last 6 months return a not found error.
	GetExchange(ctx context.Context, exchangeParams *entities.ExchangeGetParams) (entities.ExchangeTransaction, error)
}
