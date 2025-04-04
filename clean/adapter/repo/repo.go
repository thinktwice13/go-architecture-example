package repo

import (
	"clean/domain/entity"
	"clean/domain/repo"
	"clean/infra/db"
	"encoding/json"
)

type DocumentRepo struct {
	db *db.Conn
}

var _ repo.DocRepo = (*DocumentRepo)(nil)

func NewDocumentRepo(db *db.Conn) *DocumentRepo { return &DocumentRepo{db} }

func (r *DocumentRepo) Save(doc entity.Document) error {
	bytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	r.db.Set(doc.ID, bytes)
	return nil
}

func (r *DocumentRepo) FindByID(id string) (*entity.Document, error) {
	v, ok := r.db.Get(id)
	if !ok {
		return nil, nil
	}

	var doc entity.Document
	if err := json.Unmarshal(v, &doc); err != nil {
		return nil, err
	}

	return &doc, nil
}
