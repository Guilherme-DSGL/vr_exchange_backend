package services

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	irepo "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/repositories"
)

type TransactionService struct {
	transactionRepo irepo.ITransactionRepository
}

func NewTransactionService(transactionRepo irepo.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (ts *TransactionService) Save(ctx context.Context, transaction *entities.Transaction) error {
	existedTransaction, _ := ts.GetByUId(ctx, transaction.ID)

	if existedTransaction != (entities.Transaction{}) {
		return domain.ErrConflict
	}
	return ts.transactionRepo.Save(ctx, transaction)
}

func (ts *TransactionService) Fetch(ctx context.Context, cursor string, limit int64) ([]entities.Transaction, string, error) {
	transactions, nextCursor, err := ts.transactionRepo.Fetch(ctx, cursor, limit)

	if err != nil {
		return nil, "", err
	}
	return transactions, nextCursor, nil
}

func (ts *TransactionService) GetByUId(ctx context.Context, uid string) (entities.Transaction, error) {
	resptransaction, err := ts.transactionRepo.GetByUId(ctx, uid)
	if err != nil {
		return entities.Transaction{}, err
	}

	return resptransaction, nil
}

func (ts *TransactionService) Update(ctx context.Context, ar *entities.Transaction) error {
	return ts.transactionRepo.Update(ctx, ar)
}

func (ts *TransactionService) Delete(ctx context.Context, uid string) error {
	existedTransaction, err := ts.transactionRepo.GetByUId(ctx, uid)
	if err != nil {
		return err
	}
	if existedTransaction == (entities.Transaction{}) {
		return domain.ErrNotFound
	}
	return ts.transactionRepo.Delete(ctx, uid)
}
