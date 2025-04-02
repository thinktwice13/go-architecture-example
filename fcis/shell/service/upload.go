package service

import (
	"fcis/core/document"
	"fcis/shell/event"
	"fcis/shell/repo"
	"time"
)

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// UploadService handles document upload operations
// Part of the Imperative Shell - orchestrates operations and manages side effects
type UploadService struct {
	store *repo.DocumentRepo
	pub   event.Publisher
}

// NewUploadService creates a new upload service
func NewUploadService(store *repo.DocumentRepo, publisher event.Publisher) *UploadService {
	return &UploadService{store: store, pub: publisher}
}

// UploadDocument handles the document upload process
func (s *UploadService) UploadDocument(doc document.Document) error {
	// Validate document using pure function from core
	validationResult := document.ValidateForUpload(doc)
	if !validationResult.IsValid {
		return &ValidationError{Message: validationResult.Errors[0]}
	}
	// Extract metadata using pure function from core
	// Set document metadata using pure function
	// Set timestamps and status
	doc.CreatedAt = time.Now().UTC()
	doc = document.SetStatus(doc, "uploaded")

	// Save document to database (side effect)
	_ = s.store.Save(doc)

	// Publish event (side effect)
	_ = s.pub.Publish(document.Event{
		Document:  doc,
		EventType: document.EventTypeUploaded,
		Timestamp: time.Now().UTC(),
	})

	return nil
}
