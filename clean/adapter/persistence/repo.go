package persistence

import (
	"clean/domain/entity"
	"clean/infra/db"
	"clean/usecase/document"
	"fmt"
	"time"
)

// DocumentRepo implements the document repo interface
type DocumentRepo struct {
	db *db.Conn
}

var _ document.Repo = (*DocumentRepo)(nil)

// NewDocumentRepo creates a new document repo
func NewDocumentRepo(db *db.Conn) *DocumentRepo {
	return &DocumentRepo{
		db: db,
	}
}

// Save stores a document
func (r *DocumentRepo) Save(_ entity.Document) error {
	fmt.Print("Saving document to database")
	return nil
}

// FindByID retrieves a document by ID
func (r *DocumentRepo) FindByID(id int64) (*entity.Document, error) {
	fmt.Print("Getting document from database")
	doc := &entity.Document{
		ID:        id,
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}
	return doc, nil
}
