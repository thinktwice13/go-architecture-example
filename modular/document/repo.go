package document

import (
	"modular/storage"
	"strconv"
)

type Repository interface {
	Save(document Document) error
}

type InMemoryRepository struct {
	store *storage.DB
}

func (r *InMemoryRepository) Save(document Document) error {
	return r.store.Set(strconv.FormatInt(document.ID, 10), document)
}
