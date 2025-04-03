package errs

import (
	"fmt"
	"hexagonal/port/output"
)

type Handler struct{}

var _ output.ErrorHandler = (*Handler)(nil)

func NewHandler() output.ErrorHandler {
	return &Handler{}
}

func (*Handler) HandleError(err error) error {
	// Detect error here
	fmt.Println(err.Error())
	return nil
}
