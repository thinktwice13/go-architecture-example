package output

// ErrorHandler defines the output port for error handling
type ErrorHandler interface {
	HandleError(error) error
}
