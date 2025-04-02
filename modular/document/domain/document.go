package domain

import (
	"modular/common/domain"
	"time"
)

// Document represents a document in the system
type Document struct {
	ID        int64
	Name      string
	Status    string
	CreatedAt time.Time
}

// NewDocument creates a new document
func NewDocument(name string) *Document {
	now := time.Now().UTC()
	return &Document{
		Name:      name,
		Status:    "new",
		CreatedAt: now,
	}
}

// MarkAsUploaded marks the document as uploaded
func (d *Document) MarkAsUploaded() {
	d.Status = "uploaded"
}

// Validate validates the document
func (d *Document) Validate() error {
	if d.Name == "" {
		return NewValidationError("document name cannot be empty")
	}

	return nil
}

func NewValidationError(message string) error {
	return domain.NewValidationError(message)
}
