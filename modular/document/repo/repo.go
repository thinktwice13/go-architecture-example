package repo

import (
	"encoding/json"
	"fmt"
	"modular/common/infra/db"
	"modular/document/domain"
	"time"
)

type Repo struct {
	sb *db.Conn
}

var _ domain.DocumentRepo = (*Repo)(nil)

func NewRepo(sb *db.Conn) *Repo {
	return &Repo{sb}
}

func (r *Repo) Save(doc domain.Document) error {
	if bytes, err := json.Marshal(doc); err != nil {
		return err
	} else {
		r.sb.Set(doc.ID, bytes)
	}
	return nil
}

func (r *Repo) FindByID(id string) (*domain.Document, error) {
	fmt.Println("Getting document from database")
	return &domain.Document{
		ID:        id,
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	}, nil
}
