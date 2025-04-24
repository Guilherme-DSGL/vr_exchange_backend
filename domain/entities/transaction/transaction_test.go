package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShoudReturnAValidTranscation(t *testing.T) {
	transaction, err := NewTransaction(
		uuid.New(),
		"A valid description",
		time.Now(),
		2,
	)

	assert.Empty(t, err)
	assert.NotEmpty(t, transaction)
}

func TestShoudReturnErrorWhenValueIsNegative(t *testing.T) {
	transaction, err := NewTransaction(
		uuid.New(),
		"A valid description",
		time.Now(),
		-12,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, ErrTransactionInvalidValue, err)
}

func TestShoudReturnErrorWhenDescriptionIsEmpty(t *testing.T) {
	transaction, err := NewTransaction(
		uuid.New(),
		"",
		time.Now(),
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, ErrTransactionEmptyDescription, err)

}

func TestShoudReturnErrorWhenDescriptionIsLongerThan50Characters(t *testing.T) {
	invalidDescription := strings.Repeat("a", 51)
	transaction, err := NewTransaction(
		uuid.New(),
		invalidDescription,
		time.Now(),
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, ErrTransactionTooLongDescription, err)
}

func TestShoudReturnErrorWhenIdIsEmpty(t *testing.T) {
	transaction, err := NewTransaction(
		uuid.Nil,
		"",
		time.Time{},
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, ErrTransactionEmptyID, err)
}

func TestShoudReturnErrorWhenDateIsValid(t *testing.T) {
	transaction, err := NewTransaction(
		uuid.New(),
		"A valid description",
		time.Time{},
		20,
	)

	assert.Empty(t, transaction)
	assert.NotEmpty(t, err)
	assert.Equal(t, ErrTransactionInvalidDate, err)
}
