package domain

// ErrInvalidDocument represents a domain validation error
// Domain errors defined in the core and can be used by any layer
type ErrInvalidDocument string

func (e ErrInvalidDocument) Error() string {
	return string(e)
}

// ErrDocumentNotFound is returned when a document cannot be found
type ErrDocumentNotFound struct {
	ID string
}

func (e ErrDocumentNotFound) Error() string {
	return "document not found: " + e.ID
}
