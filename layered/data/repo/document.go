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

func (*Document) FindByID(_ string) (*model.Document, error) {
	fmt.Println("Finding document in the database")
	return &model.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}, nil
}
