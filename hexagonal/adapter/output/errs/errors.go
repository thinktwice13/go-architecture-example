package errs

import (
	"fmt"
	"hex/port/output"
)

// Handler implements the Handler output port
type Handler struct{}

var _ output.ErrorHandler = (*Handler)(nil)

// NewErrorHandler creates a new error handler with logging
func NewErrorHandler() output.ErrorHandler {
	return &Handler{}
}

// HandleError implements the Handler interface
func (*Handler) HandleError(err error) error {
	// Detect error here
	fmt.Println(err.Error())
	return nil
}
