package persistence

import (
	"clean/domain/entity"
	"clean/domain/repo"
	"clean/infra/db"
	"fmt"
	"time"
)

type DocumentRepo struct {
	db *db.Conn
}

var _ repo.DocRepo = (*DocumentRepo)(nil)

func NewDocumentRepo(db *db.Conn) *DocumentRepo { return &DocumentRepo{db} }

func (*DocumentRepo) Save(entity.Document) error {
	fmt.Print("Saving document to database")
	return nil
}

func (*DocumentRepo) FindByID(id int64) (*entity.Document, error) {
	fmt.Print("Getting document from database")
	return &entity.Document{
		ID:        id,
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}, nil
}
