package domain_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShoudReturnAValidTranscation(t *testing.T) {
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		"A valid description",
		time.Now(),
		2,
	)

	assert.Empty(t, err)
	assert.NotEmpty(t, transaction)
}

func TestShoudReturnErrorWhenValueIsNegative(t *testing.T) {
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		"A valid description",
		time.Now(),
		-12,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, domain.ErrTransactionInvalidValue, err)
}

func TestShoudReturnErrorWhenValueIsInvalidFormat(t *testing.T) {
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		"A valid description",
		time.Now(),
		20,
	)

	assert.Empty(t, err)
	assert.NotEmpty(t, transaction)
	assert.Equal(t, 20.0, transaction.Value)
}

func TestShoudReturnErrorWhenDescriptionIsEmpty(t *testing.T) {
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		"",
		time.Now(),
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, domain.ErrTransactionEmptyDescription, err)

}

func TestShoudReturnErrorWhenDescriptionIsLongerThan50Characters(t *testing.T) {
	invalidDescription := strings.Repeat("a", 51)
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		invalidDescription,
		time.Now(),
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, domain.ErrTransactionTooLongDescription, err)
}

func TestShoudReturnErrorWhenIdIsEmpty(t *testing.T) {
	transaction, err := domain.NewTransaction(
		"",
		"",
		time.Time{},
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, domain.ErrTransactionEmptyID, err)
}

func TestShoudReturnErrorWhenDateIsValid(t *testing.T) {
	transaction, err := domain.NewTransaction(
		uuid.New().String(),
		"A valid description",
		time.Time{},
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, domain.ErrTransactionInvalidDate, err)
}
