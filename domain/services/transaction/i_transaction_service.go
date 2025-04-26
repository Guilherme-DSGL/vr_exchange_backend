package services

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
)

//go:generate mockery --name ITransactionService
type ITransactionService interface {
	Save(ctx context.Context, t *entities.Transaction) error
	Fetch(ctx context.Context, cursor string, limit int64) ([]entities.Transaction, string, error)
	GetById(ctx context.Context, id string) (entities.Transaction, error)
	Update(ctx context.Context, t *entities.Transaction) error
	Delete(ctx context.Context, id string) error
}
