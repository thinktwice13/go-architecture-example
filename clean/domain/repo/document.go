package repo

import (
	"clean/domain/entity"
)

// DocRepo defines the entity-level interface required for document persistence
// This is shared between multiple use cases that operate on the same entity
type DocRepo interface {
	Save(doc entity.Document) error
	FindByID(id string) (*entity.Document, error)
}
