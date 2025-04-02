package persistence

import (
	"clean/adapter/logging"
	"clean/domain/entity"
	"clean/infra/storage"
	"time"
)

// DocumentRepo implements the document repo interface
type DocumentRepo struct {
	db     *storage.DB
	logger logging.Logger
}

// NewDocumentRepo creates a new document repo
func NewDocumentRepo(db *storage.DB, logger logging.Logger) *DocumentRepo {
	return &DocumentRepo{
		db:     db,
		logger: logger,
	}
}

// Save stores a document
func (r *DocumentRepo) Save(_ entity.Document) error {
	r.logger.Log("Saving document")
	return nil
}

// FindByID retrieves a document by ID
func (r *DocumentRepo) FindByID(id int64) (*entity.Document, error) {
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
func (r *DocumentRepo) Update(_ entity.Document) error {
	r.logger.Log("Updating document")
	return nil
}
