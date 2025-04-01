package entity

import (
	"fmt"
)

// ErrNotFound represents a not found error
type ErrNotFound struct {
	ID  string
	Msg string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("not found: %s (%s)", e.Msg, e.ID)
}

// ErrInvalidDocument represents a validation error
type ErrInvalidDocument struct {
	Field   string
	Message string
}

func (e ErrInvalidDocument) Error() string {
	return e.Message
}
