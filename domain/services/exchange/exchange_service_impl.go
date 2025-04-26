package services

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	irepo "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/repositories"
)

type ExchangeService struct {
	exchangeRepo    irepo.IExchangeRepository
	transactionRepo irepo.ITransactionRepository
}

func NewExchangeService(exchangeRepo irepo.IExchangeRepository, transactionRepo irepo.ITransactionRepository) *ExchangeService {
	return &ExchangeService{
		transactionRepo: transactionRepo,
		exchangeRepo:    exchangeRepo,
	}
}

func (es *ExchangeService) GetExchange(ctx context.Context, exchangeParams *entities.ExchangeGetParams) (entities.ExchangeTransaction, error) {
	existedTransaction, err := es.transactionRepo.GetById(ctx, exchangeParams.UIdTranscation)

	if existedTransaction == (entities.Transaction{}) {
		return entities.ExchangeTransaction{}, err
	}

	params := &entities.ExchangeRequestParams{
		CountryCurrency: exchangeParams.CountryCurrency,
		Value:           existedTransaction.Value,
		Date:            existedTransaction.Date,
	}

	exchangeTransaction, err := es.exchangeRepo.GetExchange(ctx, params)

	if err != nil {
		return entities.ExchangeTransaction{}, err
	}
	return exchangeTransaction, nil
}
