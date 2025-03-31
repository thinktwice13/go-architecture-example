package output

import (
	"hex/core/domain"
	"hex/ports/output"
)

var _ output.Storage = (*DB)(nil)

type DB struct{}

func (db *DB) Save(doc domain.Document) error {
	// Implement the logic to save the document to the database
	// This is a placeholder implementation
	return nil
}
