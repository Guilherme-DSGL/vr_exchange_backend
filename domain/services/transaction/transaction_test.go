package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	repo_mocks "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/repositories/mocks"
	services "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/transaction"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockRepo := new(repo_mocks.ITransactionRepository)
	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}

	transactionService := services.NewTransactionService(mockRepo)

	mockListTransactions := make([]entities.Transaction, 0)
	mockListTransactions = append(mockListTransactions, mockTransaction)

	t.Run("Should fetch successfully", func(t *testing.T) {
		mockRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListTransactions, "next-cursor", nil).Once()

		response, nextCursor, err := transactionService.Fetch(context.TODO(), "12", 10)

		assert.Equal(t, "next-cursor", nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, response, len(mockListTransactions))

		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when internal server error occurs", func(t *testing.T) {
		mockRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(nil, "", domain.ErrInternalServerError).Once()

		list, nextCursor, err := transactionService.Fetch(context.TODO(), "12", 10)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)

		mockRepo.AssertExpectations(t)
	})
}

func TestSave(t *testing.T) {
	mockRepo := new(repo_mocks.ITransactionRepository)
	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}

	transactionService := services.NewTransactionService(mockRepo)

	t.Run("Should save successfully", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrNotFound).Once()
		mockRepo.On("Save", mock.Anything,
			mock.AnythingOfType("*entities.Transaction")).Return(nil).Once()

		err := transactionService.Save(context.TODO(), &mockTransaction)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Should return conflict error when exits the same transaction", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()

		err := transactionService.Save(context.TODO(), &mockTransaction)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrConflict, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when internal server error occurs ", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrNotFound).Once()
		mockRepo.On("Save", mock.Anything,
			mock.AnythingOfType("*entities.Transaction")).Return(domain.ErrInternalServerError).Once()

		err := transactionService.Save(context.TODO(), &mockTransaction)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByUId(t *testing.T) {
	mockRepo := new(repo_mocks.ITransactionRepository)
	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}
	transactionService := services.NewTransactionService(mockRepo)

	t.Run("Should get successfully", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()

		response, err := transactionService.GetById(context.TODO(), mockTransaction.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockTransaction.ID, response.ID)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Should return a error when internal server error occurs ", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrInternalServerError).Once()

		response, err := transactionService.GetById(context.TODO(), mockTransaction.ID)

		assert.Error(t, err)
		assert.Empty(t, response)
		mockRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockRepo := new(repo_mocks.ITransactionRepository)
	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}
	transactionService := services.NewTransactionService(mockRepo)

	t.Run("Should delete successfully", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()
		mockRepo.On("Delete", mock.Anything,
			mock.AnythingOfType("string")).Return(nil).Once()

		err := transactionService.Delete(context.TODO(), mockTransaction.ID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Should return not found error when no exits the transaction", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(entities.Transaction{}, domain.ErrNotFound).Once()

		err := transactionService.Delete(context.TODO(), mockTransaction.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrNotFound, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when internal server error occurs ", func(t *testing.T) {
		mockRepo.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()
		mockRepo.On("Delete", mock.Anything,
			mock.AnythingOfType("string")).Return(domain.ErrInternalServerError).Once()

		err := transactionService.Delete(context.TODO(), mockTransaction.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInternalServerError, domain.ErrInternalServerError)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockRepo := new(repo_mocks.ITransactionRepository)
	mockTransaction := entities.Transaction{
		ID:          uuid.NewString(),
		Description: "Descrição valida",
		Date:        time.Now(),
		Value:       4.0,
	}
	transactionService := services.NewTransactionService(mockRepo)

	t.Run("Should update successfully", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything,
			mock.AnythingOfType("*entities.Transaction")).Return(nil).Once()

		err := transactionService.Update(context.TODO(), &mockTransaction)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return a error when update fails", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything,
			mock.AnythingOfType("*entities.Transaction")).Return(domain.ErrInternalServerError).Once()

		err := transactionService.Update(context.TODO(), &mockTransaction)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
