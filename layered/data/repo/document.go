package repo

import (
	"encoding/json"
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

func (d *Document) Save(doc model.Document) error {
	bytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	d.db.Set(doc.ID, bytes)
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
