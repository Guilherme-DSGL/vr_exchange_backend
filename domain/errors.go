package domain

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found error")

	// Transaction
	ErrTransactionEmptyID            = errors.New("id should not be null")
	ErrTransactionEmptyDescription   = errors.New("description should not be empty or null")
	ErrTransactionTooLongDescription = errors.New("description should not be longer than 50 characters")
	ErrTransactionInvalidDate        = errors.New("invalid date")
	ErrTransactionInvalidValue       = errors.New("value should be a positive number")
)
