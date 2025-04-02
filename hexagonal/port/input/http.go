package input

import "hexagonal/core/domain"

type Http interface {
	Upload(doc domain.Document) error
	Find(id int64) (*domain.Document, error)
}
