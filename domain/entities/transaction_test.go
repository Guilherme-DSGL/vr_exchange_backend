package entities_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {

	t.Run("Shoud return a valid transcation", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			"A valid description",
			time.Now(),
			2,
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, transaction)
	})

	t.Run("Shoud return error when value is negative", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			"A valid description",
			time.Now(),
			-12,
		)

		assert.Empty(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTransactionInvalidValue, err)
	})

	t.Run("Shoud return error when value is invalid format", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			"A valid description",
			time.Now(),
			20,
		)

		assert.Empty(t, err)
		assert.NotEmpty(t, transaction)
		assert.Equal(t, 20.0, transaction.Value)
	})

	t.Run("Shoud return error when description is empty", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			"",
			time.Now(),
			20,
		)

		assert.Empty(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTransactionEmptyDescription, err)

	})

	t.Run("Shoud return error when description is longer than 50 characters", func(t *testing.T) {
		invalidDescription := strings.Repeat("a", 51)
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			invalidDescription,
			time.Now(),
			20,
		)

		assert.Empty(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTransactionTooLongDescription, err)
	})

	t.Run("Shoud return error when uid is empty", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			"",
			"",
			time.Time{},
			20,
		)

		assert.Empty(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTransactionEmptyID, err)
	})

	t.Run("Shoud return error when date is valid", func(t *testing.T) {
		transaction, err := entities.NewTransaction(
			uuid.NewString(),
			"A valid description",
			time.Time{},
			20,
		)

		assert.Empty(t, transaction)
		assert.Error(t, err)
		assert.Equal(t, domain.ErrTransactionInvalidDate, err)
	})
}
