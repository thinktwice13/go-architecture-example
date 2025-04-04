package output

import (
	"hexagonal/core/domain"
)

// DocumentRepo defines the output port for document storage
// This interface is implemented by secondary adapters and used by core services
type DocumentRepo interface {
	Save(doc domain.Document) error
	FindByID(id string) (*domain.Document, error)
}
