package repo

import (
	"fmt"
	"layered/data/db"
	"layered/data/model"
	"time"
)

type Document struct {
	db *db.Conn
}

func NewDocumentRepo(db *db.Conn) *Document {
	return &Document{db: db}
}

func (*Document) Save(_ model.Document) error {
	fmt.Println("Saving document to the database")
	return nil
}

func (*Document) FindByID(_ int64) (*model.Document, error) {
	fmt.Println("Finding document in the database")
	return &model.Document{
		ID:        1,
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}, nil
}
