package repo

import (
	"fmt"
	"hexagonal/adapter/output/db"
	"hexagonal/core/domain"
	"hexagonal/port/output"
	"time"
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

func (*Repo) Save(_ domain.Document) error {
	fmt.Println("Saving document...")
	return nil
}

func (*Repo) FindByID(_ int64) (*domain.Document, error) {
	fmt.Println("Getting document...")
	return &domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "sample.txt",
		CreatedAt: time.Now(),
	}, nil
}
