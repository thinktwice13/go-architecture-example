package document

import (
	"clean/domain/entity"
	"clean/domain/event"
)

// Repository defines the entity-level interface required for document persistence
// This is shared between multiple use cases that operate on the same entity
type Repository interface {
	Save(doc entity.Document) error
	FindByID(id int64) (*entity.Document, error)
	Update(doc entity.Document) error
}

// EventPublisher defines the interface for publishing domain pub
type EventPublisher interface {
	Publish(event event.DocumentEvent) error
}
