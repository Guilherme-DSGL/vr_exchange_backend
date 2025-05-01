package services

import (
	"context"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	irepo "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/repositories"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/utils"
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
	existedTransaction, err := es.transactionRepo.GetById(ctx, exchangeParams.IdTranscation)

	if err == domain.ErrNotFound {
		return entities.ExchangeTransaction{}, domain.ErrBadParamInput
	}

	if err != nil {
		return entities.ExchangeTransaction{}, domain.ErrInternalServerError
	}

	endDate := existedTransaction.Date.Format(utils.DateFormat)
	lastSixMonths := -6
	startDate := existedTransaction.Date.AddDate(0, lastSixMonths, 0).Format(utils.DateFormat)
	params := &entities.ExchangeRequestParams{
		CountryCurrency: exchangeParams.CountryCurrency,
		Value:           existedTransaction.Value,
		StartDate:       startDate,
		EndDate:         endDate,
	}

	exchangeTransaction, err := es.exchangeRepo.GetExchange(ctx, params)
	exchangeTransaction.Transaction = existedTransaction

	if err == domain.ErrNotFound {
		return entities.ExchangeTransaction{}, err
	}

	if err != nil {
		return entities.ExchangeTransaction{}, domain.ErrInternalServerError
	}
	return exchangeTransaction, nil
}
