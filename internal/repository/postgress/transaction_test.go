package postgress_test

import (
	"context"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/internal/repository"

	postgressRepo "github.com/Guilherme-DSGL/purchase_transaction_backend/internal/repository/postgress"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPostgress(t *testing.T) {
	var colums = []string{"id", "description", "value", "date", "updated_at", "created_at"}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockTransactions := []entities.Transaction{
		{
			ID:          uuid.NewString(),
			Description: "Lego star wars",
			Value:       20.00,
			Date:        time.Now(),
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.NewString(),
			Description: "Lego star wars",
			Value:       20.00,
			Date:        time.Now(),
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		},
	}

	t.Run("Should fecth query return transactions", func(t *testing.T) {

		rows := sqlmock.NewRows(colums).
			AddRow(mockTransactions[0].ID, mockTransactions[0].Description, mockTransactions[0].Date,
				mockTransactions[0].Value, mockTransactions[0].UpdatedAt, mockTransactions[0].CreatedAt).
			AddRow(mockTransactions[1].ID, mockTransactions[1].Description, mockTransactions[1].Date,
				mockTransactions[1].Value, mockTransactions[1].UpdatedAt, mockTransactions[1].CreatedAt)
		query := "SELECT id, description, date, value, updated_at, created_at FROM transaction WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

		mock.ExpectQuery(query).WillReturnRows(rows)

		a := postgressRepo.NewTransactionRepository(db)

		cursor := repository.EncodeCursor(mockTransactions[1].CreatedAt)
		num := int64(2)
		list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, 2)
	})

	t.Run("Should get transaction by id", func(t *testing.T) {
		mockTransaction := mockTransactions[0]
		rows := sqlmock.NewRows(colums).
			AddRow(mockTransaction.ID, mockTransaction.Description, mockTransaction.Date,
				mockTransaction.Value, mockTransaction.UpdatedAt, mockTransaction.CreatedAt)

		query := "SELECT id, description, date, value, updated_at, created_at FROM transaction WHERE ID = \\?"

		mock.ExpectQuery(query).WillReturnRows(rows)
		tr := postgressRepo.NewTransactionRepository(db)

		id := mockTransactions[0].ID
		println(id)
		anTransaction, err := tr.GetById(context.TODO(), id)
		println(anTransaction.ID)
		assert.NoError(t, err)
		assert.NotNil(t, anTransaction)
	})

	t.Run("Should save a transaction", func(t *testing.T) {
		tr := &entities.Transaction{
			ID:          uuid.NewString(),
			Description: "valid description",
			Date:        time.Now(),
			Value:       20.0,
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		}

		query := "INSERT  transaction SET id=\\?, description=\\?, date=\\?, value=\\?, updated_at=\\?, created_at=\\?"
		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(tr.ID, tr.Description, tr.Date, tr.Value, tr.UpdatedAt, tr.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

		a := postgressRepo.NewTransactionRepository(db)

		err = a.Save(context.TODO(), tr)
		assert.NoError(t, err)
	})

	t.Run("Should delete a transaction", func(t *testing.T) {

		query := "DELETE FROM transaction WHERE id = \\?"

		prep := mock.ExpectPrepare(query)

		id := mockTransactions[0].ID
		prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(12, 1))

		a := postgressRepo.NewTransactionRepository(db)

		err = a.Delete(context.TODO(), id)
		assert.NoError(t, err)
	})

	t.Run("Should update a transaction", func(t *testing.T) {

		tr := &entities.Transaction{
			ID:          uuid.NewString(),
			Description: "valid description",
			Date:        time.Now(),
			Value:       20.0,
			UpdatedAt:   time.Now(),
			CreatedAt:   time.Now(),
		}

		query := "UPDATE transaction SET description=\\?, date=\\?, value=\\?, updated_at=\\? WHERE ID = \\?"

		prep := mock.ExpectPrepare(query)
		prep.ExpectExec().WithArgs(tr.Description, tr.Date, tr.Value, tr.UpdatedAt, tr.ID).WillReturnResult(sqlmock.NewResult(12, 1))

		a := postgressRepo.NewTransactionRepository(db)

		err = a.Update(context.TODO(), tr)
		assert.NoError(t, err)
	})
}
