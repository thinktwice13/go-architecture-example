package retrieve

import (
	"hex/core/domain"
	"vert/shared/db"
)

type Repository struct {
	db *db.InMemoryDB
}

func (r *Repository) FindByID(id string) (domain.Document, error) {
	doc, err := r.db.Get(id)
	if err != nil {
		return domain.Document{}, err
	}
	return doc.(domain.Document), nil
}
