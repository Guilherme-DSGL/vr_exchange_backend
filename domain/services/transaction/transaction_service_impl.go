package services

import (
	"context"
	"time"

	"github.com/Guilherme-DSGL/vr_exchange_backend/domain"
	"github.com/Guilherme-DSGL/vr_exchange_backend/domain/entities"
	irepo "github.com/Guilherme-DSGL/vr_exchange_backend/domain/repositories"
	"github.com/Guilherme-DSGL/vr_exchange_backend/utils"
	"github.com/google/uuid"
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
	transaction.ID = uuid.NewString()
	existedTransaction, _ := ts.GetById(ctx, transaction.ID)

	if existedTransaction != (entities.Transaction{}) {
		return domain.ErrConflict
	}
	transaction.Value = utils.RoundValue2DecimalPlaces(transaction.Value)
	transaction.CreatedAt = time.Now()
	err := ts.transactionRepo.Save(ctx, transaction)

	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (ts *TransactionService) Fetch(ctx context.Context, cursor string, limit int64) ([]entities.Transaction, string, error) {
	transactions, nextCursor, err := ts.transactionRepo.Fetch(ctx, cursor, limit)
	if err == domain.ErrBadParamInput {
		return nil, "", err
	}
	if err != nil {
		return nil, "", domain.ErrInternalServerError
	}
	return transactions, nextCursor, nil
}

func (ts *TransactionService) GetById(ctx context.Context, id string) (entities.Transaction, error) {
	resptransaction, err := ts.transactionRepo.GetById(ctx, id)
	if err == domain.ErrNotFound {
		return entities.Transaction{}, err
	}

	if err != nil {
		return entities.Transaction{}, domain.ErrInternalServerError
	}

	return resptransaction, nil
}

func (ts *TransactionService) Update(ctx context.Context, transaction *entities.Transaction) error {
	transaction.Value = utils.RoundValue2DecimalPlaces(transaction.Value)
	transaction.UpdatedAt = time.Now()
	return ts.transactionRepo.Update(ctx, transaction)
}

func (ts *TransactionService) Delete(ctx context.Context, id string) error {
	existedTransaction, err := ts.transactionRepo.GetById(ctx, id)
	if existedTransaction == (entities.Transaction{}) {
		return domain.ErrNotFound
	}
	if err != nil {
		return domain.ErrInternalServerError
	}

	err = ts.transactionRepo.Delete(ctx, id)

	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
