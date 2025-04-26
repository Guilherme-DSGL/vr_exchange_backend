package entities_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestNewTransaction(t *testing.T) {
	validate := validator.New()

	t.Run("Shoud return a valid transcation", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       2,
		}
		err := validate.Struct(transaction)
		assert.NoError(t, err)
		assert.NotEmpty(t, transaction)
	})

	t.Run("Shoud return error when value is negative", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       -12,
		}
		err := validate.Struct(transaction)
		assert.Error(t, err)
	})

	t.Run("Shoud return error when value is invalid format", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Now(),
			Value:       20.56711312123,
		}

		transaction.Value = utils.RoundValue2DecimalPlaces(transaction.Value)
		err := validate.Struct(transaction)
		assert.Empty(t, err)
		assert.NotEmpty(t, transaction)
		assert.Equal(t, "20.570000", fmt.Sprintf("%f", transaction.Value))
	})

	t.Run("Shoud return error when description is empty", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "",
			Date:        time.Now(),
			Value:       20,
		}
		err := validate.Struct(transaction)
		assert.Error(t, err)
	})

	t.Run("Shoud return error when description is longer than 50 characters", func(t *testing.T) {
		invalidDescription := strings.Repeat("a", 51)
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: invalidDescription,
			Date:        time.Now(),
			Value:       20,
		}
		err := validate.Struct(transaction)

		assert.Error(t, err)
	})

	t.Run("Shoud return error when id is empty", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          "",
			Description: "",
			Date:        time.Now(),
			Value:       20,
		}
		err := validate.Struct(transaction)
		assert.Error(t, err)
	})

	t.Run("Shoud return error when date is invalid", func(t *testing.T) {
		transaction := entities.Transaction{
			ID:          uuid.NewString(),
			Description: "A valid description",
			Date:        time.Time{},
			Value:       20,
		}

		err := validate.Struct(transaction)
		assert.Error(t, err)
	})
}
