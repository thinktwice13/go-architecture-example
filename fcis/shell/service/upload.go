package service

import (
	"fcis/core/document"
	"fcis/shell/event"
	"fcis/shell/repo"
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

// UploadDocument handles the document upload process.
// While not very idiomatic for Go, this is a good example of the shell using the pure functional core.
func (s *UploadService) UploadDocument(doc document.Document) error {
	// Validate document using function from core (pure), that maybe returns an error (side effect)
	validationResult := document.ValidateForUpload(doc)
	if !validationResult.IsValid {
		return &ValidationError{Message: validationResult.Errors[0]}
	}

	// Use functional core to update (pure) that returns a new document, then save it to repo (side effect)
	uploaded := document.UpdateStatus(doc, "uploaded")
	_ = s.store.Save(doc)

	// Use functional core to create event (pure), then publish it (side effect)
	return s.pub.Publish(document.NewUploadedEvent(uploaded))
}
