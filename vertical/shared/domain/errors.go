package domain

type ErrInvalidDocument string

func (e ErrInvalidDocument) Error() string {
	return string(e)
}

type ErrDocumentNotFound struct {
	ID string
}

func (e ErrDocumentNotFound) Error() string {
	return "document not found: " + e.ID
}
