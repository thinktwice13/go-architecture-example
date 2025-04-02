package errs

import (
	"fmt"
	"hexagonal/port/output"
)

// Handler implements the Handler output port
type Handler struct{}

var _ output.ErrorHandler = (*Handler)(nil)

// NewHandler creates a new error handler with logging
func NewHandler() output.ErrorHandler {
	return &Handler{}
}

// HandleError implements the Handler interface
func (*Handler) HandleError(err error) error {
	// Detect error here
	fmt.Println(err.Error())
	return nil
}
