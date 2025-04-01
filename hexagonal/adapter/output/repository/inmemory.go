package repository

import (
	"fmt"
	"hex/adapter/output/db"
	"hex/core/domain"
	"hex/port/output"
	"time"
)

// InMemory implements the DocumentRepository output port
// This is a secondary adapter that is driven by the application
type InMemory struct {
	db *db.DB
}

var _ output.DocumentRepository = (*InMemory)(nil)

// NewInMemoryRepository creates a new Postgres-based document repository
func NewInMemoryRepository(db *db.DB) output.DocumentRepository {
	return &InMemory{db}
}

func (m *InMemory) Save(_ domain.Document) error {
	fmt.Println("Saving document to in-memory database")
	return nil
}

func (m *InMemory) FindByID(_ int64) (*domain.Document, error) {
	fmt.Println("Getting document from in-memory database")
	return &domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "sample.txt",
		CreatedAt: time.Now(),
	}, nil
}
