package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	repo_mocks "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/repositories/mocks"
	services "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/exchange"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExchange(t *testing.T) {
	mockExchangeRepo := new(repo_mocks.IExchangeRepository)
	mockTransactionRepo := new(repo_mocks.ITransactionRepository)
	mockCountryCurrency := entities.CountryCurreny{
		Country:  "Brazil",
		Currency: "Real",
	}

	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}
	mockGetParams := entities.ExchangeGetParams{
		IdTranscation:   mockTransaction.ID,
		CountryCurrency: mockCountryCurrency,
	}
	mockExchangeTransaction := entities.ExchangeTransaction{
		Transaction:    mockTransaction,
		CountryCurreny: mockCountryCurrency,
		EfectiveDate:   time.Now(),
		Rate:           1.0,
		ConvertedValue: 20.0,
	}

	exchangeService := services.NewExchangeService(
		mockExchangeRepo,
		mockTransactionRepo,
	)

	t.Run("Should exchange successfully", func(t *testing.T) {
		mockTransactionRepo.On("GetById",
			mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()
		mockExchangeRepo.On("GetExchange", mock.Anything,
			mock.AnythingOfType("*entities.ExchangeRequestParams"),
		).Return(mockExchangeTransaction, nil)

		exchangeTransaction, err := exchangeService.GetExchange(context.TODO(), &mockGetParams)

		assert.NoError(t, err)
		assert.NotEmpty(t, exchangeTransaction)
		mockTransactionRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when not found transaction", func(t *testing.T) {
		mockTransactionRepo.On("GetById", mock.Anything,
			mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrNotFound).Once()

		exchangeService := services.NewExchangeService(
			mockExchangeRepo,
			mockTransactionRepo,
		)

		exchangeTransaction, err := exchangeService.GetExchange(context.TODO(), &mockGetParams)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrBadParamInput, err)
		assert.Empty(t, exchangeTransaction)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when internal server erro occurs", func(t *testing.T) {
		mockTransactionRepo.On("GetById", mock.Anything,
			mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrInternalServerError).Once()
		mockExchangeRepo.On("GetExchange", mock.Anything,
			mock.AnythingOfType("*entities.ExchangeRequestParams"),
		).Return(entities.ExchangeTransaction{}, domain.ErrInternalServerError)
		exchangeService := services.NewExchangeService(
			mockExchangeRepo,
			mockTransactionRepo,
		)

		exchangeTransaction, err := exchangeService.GetExchange(context.TODO(), &mockGetParams)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrInternalServerError, err)
		assert.Empty(t, exchangeTransaction)
		mockTransactionRepo.AssertExpectations(t)
		mockExchangeRepo.AssertExpectations(t)
	})
}
