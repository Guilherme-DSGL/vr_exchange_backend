package irepo

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
)

//go:generate mockery --name ITransactionRepository
type ITransactionRepository interface {
	Save(ctx context.Context, transaction *entities.Transaction) error
	Fetch(ctx context.Context, cursor string, limit int64) ([]entities.Transaction, string, error)
	GetByUId(ctx context.Context, uid string) (entities.Transaction, error)
	Update(ctx context.Context, transaction *entities.Transaction) error
	Delete(ctx context.Context, uid string) error
}
