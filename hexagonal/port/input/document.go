package input

import (
	"hex/core/domain"
)

// DocumentService defines the input port for document operations
// This interface is implemented by the core service and used by primary adapters
type DocumentService interface {
	UploadDocument(doc domain.Document) error
	GetDocument(id int64) (*domain.Document, error)
}
