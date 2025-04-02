package repo

import (
	"fcis/core/document"
	"fcis/shell/db"
	"fmt"
	"time"
)

// DocumentRepo handles document storage operations
// Part of the Imperative Shell - manages database side effects
type DocumentRepo struct {
	db *db.Conn
}

func NewDocumentRepo(dbConn *db.Conn) *DocumentRepo {
	return &DocumentRepo{
		db: dbConn,
	}
}

func (*DocumentRepo) Save(doc document.Document) error {
	_ = document.ValidateForStorage(doc)
	fmt.Println("Saving document to database")
	return nil
}

func (*DocumentRepo) FindByID(_ int64) (*document.Document, error) {
	fmt.Println("Getting document from database")
	return &document.Document{
		ID:        time.Now().UnixNano(),
		Name:      "sample.txt",
		CreatedAt: time.Now().UTC(),
	}, nil
}
