package document

import (
	"modular/storage"
	"strconv"
)

type Repo interface {
	Save(document Document) error
}

type InMemoryRepo struct {
	store *storage.DB
}

func (r *InMemoryRepo) Save(document Document) error {
	return r.store.Set(strconv.FormatInt(document.ID, 10), document)
}
