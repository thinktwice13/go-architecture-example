package document

import (
	"clean/adapter/logging"
	"clean/domain/entity"
	"clean/domain/event"
	"context"
	"time"
)

// RetrieveUseCase implements the document retrieval use case
// Separation at use case level for distinct features
type RetrieveUseCase struct {
	repo   Repository // Same repository interface, shared across use cases
	pub    EventPublisher
	logger logging.Logger
}

// NewRetrieveUseCase creates a new retrieve use case
func NewRetrieveUseCase(repo Repository, events EventPublisher, logger logging.Logger) *RetrieveUseCase {
	return &RetrieveUseCase{
		repo:   repo,
		pub:    events,
		logger: logger,
	}
}

// Retrieve fetches a document by ID
func (uc *RetrieveUseCase) Retrieve(ctx context.Context, id int64) (*entity.Document, error) {
	doc, _ := uc.repo.FindByID(id)

	// Publish event
	_ = uc.pub.Publish(event.DocumentRetrieved{
		DocumentID: doc.ID,
		Timestamp:  time.Now(),
	})

	uc.logger.Log("Document retrieved successfully")
	return doc, nil
}
