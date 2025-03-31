package storage

import (
	"fcis/core/document"
	"fcis/shell/infra"
	"strconv"
)

type Repository interface {
	Save(document document.Document) error
}

var _ Repository = &DocumentStore{}

type DocumentStore struct {
	db *infra.DB
}

func NewDocumentStore(db *infra.DB) *DocumentStore {
	return &DocumentStore{db}
}

func (s *DocumentStore) Save(doc document.Document) error {
	s.db.Set(strconv.Itoa(int(doc.ID)), doc)
	return nil
}
