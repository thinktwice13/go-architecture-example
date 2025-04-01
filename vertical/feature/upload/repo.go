package upload

import (
	"hex/core/domain"
	"strconv"
	"vert/shared/db"
)

type Repository struct {
	db *db.DB
}

func newRepo(db *db.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(doc domain.Document) error {
	return r.db.Set(strconv.Itoa(int(doc.ID)), doc)
}
