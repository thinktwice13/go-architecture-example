package repo

import (
	"encoding/json"
	"fmt"
	"hexagonal/adapter/output/db"
	"hexagonal/core/domain"
	"hexagonal/port/output"
)

// Repo implements the DocumentRepo output port
// This is a secondary adapter that is driven by the application
type Repo struct {
	db *db.Conn
}

var _ output.DocumentRepo = (*Repo)(nil)

func NewRepo(db *db.Conn) output.DocumentRepo {
	return &Repo{db}
}

func (r *Repo) Save(doc domain.Document) error {
	bytes, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal document: %w", err)
	}

	r.db.Set(doc.ID, bytes)
	return nil
}

func (r *Repo) FindByID(id string) (*domain.Document, error) {
	v, ok := r.db.Get(id)
	if !ok {
		return nil, fmt.Errorf("document not found")
	}
	var doc domain.Document
	if err := json.Unmarshal(v, &doc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal document: %w", err)
	}

	return &doc, nil
}
