package domain

import "errors"

var (
	ErrTransactionEmptyID            = errors.New("id should not be null")
	ErrTransactionEmptyDescription   = errors.New("description should not be empty or null")
	ErrTransactionTooLongDescription = errors.New("description should not be longer than 50 characters")
	ErrTransactionInvalidDate        = errors.New("invalid date")
	ErrTransactionInvalidValue       = errors.New("value should be a positive number")
)
