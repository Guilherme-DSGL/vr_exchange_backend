package services

import (
	"context"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
)

// Represents the business logic for handling transaction operations.
//
//go:generate mockery --name ITransactionService
type ITransactionService interface {
	// Save the given transaction to transaction repository
	// Returns an error if the operation fails.
	Save(ctx context.Context, t *entities.Transaction) error
	// Retrieves a list of transactions starting from the given cursor,
	// Returns the nextCurosr or error if the operation fails
	Fetch(ctx context.Context, cursor string, limit int64) ([]entities.Transaction, string, error)
	// GetById retrieves a transaction by its unique ID.
	// Returns an error if the transaction is not found.
	GetById(ctx context.Context, id string) (entities.Transaction, error)
	// Update modifies an existing transaction in the repository.
	// Returns an error if the update fails or transaction does'nt exists.
	Update(ctx context.Context, t *entities.Transaction) error
	// Delete removes a transaction from the repository by ID.
	// Returns an error if the deletion fails or transaction does'nt exists.
	Delete(ctx context.Context, id string) error
}
