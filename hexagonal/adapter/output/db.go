package output

import (
	"hex/core/domain"
	"hex/port/output"
)

var _ output.Storage = (*DB)(nil)

type DB struct{}

func (db *DB) Save(doc domain.Document) error {
	// Implement the logic to save the document to the database
	// This is a placeholder implementation
	return nil
}

func (db *DB) FindByID(id int64) (*domain.Document, error) {
	// Implement the logic to find a document by ID in the database
	// This is a placeholder implementation
	return nil, nil
}
