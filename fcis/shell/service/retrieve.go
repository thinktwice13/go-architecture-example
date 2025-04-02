package service

import (
	"fcis/core/document"
	"fcis/shell/repo"
)

// NotFoundError represents a not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// RetrieveService handles document retrieval operations
type RetrieveService struct {
	store *repo.DocumentRepo
}

// NewRetrieveService creates a new retrieve service
func NewRetrieveService(store *repo.DocumentRepo) *RetrieveService {
	return &RetrieveService{store: store}
}

// GetDocument retrieves a document by ID
func (s *RetrieveService) GetDocument(id int64) (*document.Document, error) {
	// Retrieve document from database (side effect)
	doc, err := s.store.FindByID(id)
	if err != nil {
		if err.Error() == "document not found" {
			return doc, &NotFoundError{Message: "Document not found"}
		}
		return doc, err
	}

	return doc, nil
}
