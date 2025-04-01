package retrieve

import (
	"hex/core/domain"
	"vert/shared/db"
)

type Repository struct {
	db *db.DB
}

func (r *Repository) FindByID(id string) (*domain.Document, error) {
	v, err := r.db.Get(id)
	if err != nil {
		return nil, domain.ErrDocumentNotFound
	}

	doc, ok := v.(domain.Document)
	if !ok {
		return nil, err
	}
	return &doc, nil
}
