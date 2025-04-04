package input

import (
	"hexagonal/core/domain"
)

// DocumentService defines the input port for document operations
// This interface is implemented by the core service and used by primary adapters
type DocumentService interface {
	Upload(doc domain.Document) error
	Find(id string) (*domain.Document, error)
}
