package entity

import "errors"

var ErrNotFound = errors.New("resource not found")

// ErrInvalidDocument represents a validation error
type ErrInvalidDocument struct {
	Field   string
	Message string
}

func (e ErrInvalidDocument) Error() string {
	return e.Message
}
