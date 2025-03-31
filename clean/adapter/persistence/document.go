package persistence

import (
	"clean/entity"
	"clean/infra/storage"
	"clean/usecase/iface"
	"errors"
	"strconv"
)

var (
	_ iface.DocumentRepository = (*DocumentRepository)(nil)

	ErrNotFound = errors.New("document not found")
)

type DocumentRepository struct {
	db *storage.DB
}

func NewDocumentRepository(db *storage.DB) *DocumentRepository {
	return &DocumentRepository{db}
}

func (r *DocumentRepository) Save(doc entity.Document) error {
	return r.db.Set(strconv.FormatInt(doc.ID, 10), doc)
}

func (r *DocumentRepository) Find(id int64) (*entity.Document, error) {
	v, err := r.db.Get(strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	doc, ok := v.(entity.Document)
	if !ok {
		return nil, ErrNotFound
	} else {
		doc = v.(entity.Document)
	}
	return &doc, nil
}

func (r *DocumentRepository) UpdateStatus(id int64, status string) error {
	doc, err := r.Find(id)
	if err != nil {
		return err
	}
	if doc == nil {
		return ErrNotFound
	}
	doc.Status = status
	return r.Save(*doc)
}
