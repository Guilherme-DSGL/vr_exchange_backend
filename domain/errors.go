package domain

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found error")
	ErrConflict            = errors.New("conflict error the same identifier already exists")
)
