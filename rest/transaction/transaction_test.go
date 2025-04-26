package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	mocks_transactions "github.com/Guilherme-DSGL/purchase_transaction_backend/domain/services/transaction/mocks"
	rest "github.com/Guilherme-DSGL/purchase_transaction_backend/rest/transaction"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestTransactionHandler(t *testing.T) {
	var mockTransaction entities.Transaction
	err := faker.FakeData(&mockTransaction)
	assert.NoError(t, err)
	mockUCase := new(mocks_transactions.ITransactionService)
	validate := validator.New()

	t.Run("Should fetch return status 200", func(t *testing.T) {

		mockListTransactions := make([]entities.Transaction, 0)
		mockListTransactions = append(mockListTransactions, mockTransaction)
		num := 10
		cursor := "2"
		nextCursor := "12"
		mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListTransactions, nextCursor, nil)

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.GET, "/transaction?num=10&cursor="+cursor, strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		responseCursor := rec.Header().Get("X-Cursor")
		assert.Equal(t, nextCursor, responseCursor)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should fetch return a error", func(t *testing.T) {

		mockListTransactions := make([]entities.Transaction, 0)
		mockListTransactions = append(mockListTransactions, mockTransaction)

		mockUCase.On("Fetch", mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListTransactions, "", domain.ErrInternalServerError)

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.GET, "/transaction?num=10&cursor=0", strings.NewReader(""))
		assert.NoError(t, err)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should GetById return status 200", func(t *testing.T) {
		id := uuid.NewString()
		mockUCase.On("GetById", mock.Anything, mock.AnythingOfType("string")).Return(mockTransaction, nil).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.GET, "/transaction/:id="+id, strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should GetById return a error", func(t *testing.T) {
		id := uuid.NewString()
		mockUCase.On("GetById",
			mock.Anything,
			mock.AnythingOfType("string")).Return(mockTransaction, domain.ErrInternalServerError).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.GET, "/transaction/:id="+id, strings.NewReader(""))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should Save return status 200", func(t *testing.T) {
		mockTransaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       2,
		}
		err := validate.Struct(mockTransaction)
		assert.NoError(t, err)

		json, err := json.Marshal(mockTransaction)

		assert.NoError(t, err)

		mockUCase.On("Save", mock.Anything, mock.AnythingOfType("*entities.Transaction")).Return(nil).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.POST, "/transaction", strings.NewReader(string(json)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should Save return a error", func(t *testing.T) {
		mockTransaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       2,
		}
		err := validate.Struct(mockTransaction)
		assert.NoError(t, err)

		json, err := json.Marshal(mockTransaction)

		assert.NoError(t, err)

		mockUCase.On("Save", mock.Anything, mock.AnythingOfType("*entities.Transaction")).Return(domain.ErrInternalServerError).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.POST, "/transaction", strings.NewReader(string(json)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should Delete return status 200", func(t *testing.T) {
		id := uuid.NewString()
		mockTransaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       2,
		}
		err := validate.Struct(mockTransaction)
		assert.NoError(t, err)

		mockUCase.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.DELETE, "/transaction/:id="+id, strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Should Delete return a error", func(t *testing.T) {
		id := uuid.NewString()
		mockTransaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       2,
		}
		err := validate.Struct(mockTransaction)
		assert.NoError(t, err)

		mockUCase.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(domain.ErrInternalServerError).Once()

		e := echo.New()
		rest.NewTransactionHandler(e, mockUCase)
		req, err := http.NewRequestWithContext(context.TODO(),
			echo.DELETE, "/transaction/:id="+id, strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})

}
