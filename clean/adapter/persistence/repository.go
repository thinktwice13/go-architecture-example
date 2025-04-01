package persistence

import (
	"clean/adapter/logging"
	"clean/domain/entity"
	"clean/infra/storage"
	"time"
)

// DocumentRepository implements the document repository interface
type DocumentRepository struct {
	db     *storage.DB
	logger logging.Logger
}

// NewDocumentRepository creates a new document repository
func NewDocumentRepository(db *storage.DB, logger logging.Logger) *DocumentRepository {
	return &DocumentRepository{
		db:     db,
		logger: logger,
	}
}

// Save stores a document
func (r *DocumentRepository) Save(_ entity.Document) error {
	r.logger.Log("Saving document")
	return nil
}

// FindByID retrieves a document by ID
func (r *DocumentRepository) FindByID(id int64) (*entity.Document, error) {
	r.logger.Log("Finding document by ID")
	doc := &entity.Document{
		ID:        id,
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}
	return doc, nil
}

// Update modifies an existing document
func (r *DocumentRepository) Update(_ entity.Document) error {
	r.logger.Log("Updating document")
	return nil
}
