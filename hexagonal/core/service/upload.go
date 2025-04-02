package service

import (
	"hexagonal/core/domain"
	"hexagonal/port/input"
	"hexagonal/port/output"
	"time"
)

var (
	_ input.DocumentService = (*DocumentService)(nil)
)

// DocumentService implements the input ports (use cases) for document operations
// This is the core application service that orchestrates the domain logic
type DocumentService struct {
	repo   output.DocumentRepo
	pub    output.EventPublisher
	errors output.ErrorHandler
}

// func NewDocumentService(repo output.DocumentRepo, eb output.EventPublisher) *DocumentService {
// 	return &DocumentService{repo, eb}
// }

// NewDocumentService creates a new document service with required dependencies
func NewDocumentService(
	repo output.DocumentRepo,
	pub output.EventPublisher,
	errors output.ErrorHandler,
) *DocumentService {
	return &DocumentService{repo, pub, errors}
}

func (s *DocumentService) Upload(doc domain.Document) error {
	if err := doc.Validate(); err != nil {
		return s.errors.HandleError(err)
	}

	// Business logic...
	_ = s.repo.Save(doc)
	_ = s.pub.Publish(domain.DocumentEvent{Type: domain.EventDocumentUploaded, Timestamp: time.Now().UTC(), Document: &doc})
	return nil
}

func (s *DocumentService) Find(id int64) (*domain.Document, error) {
	return s.repo.FindByID(id)
}
