package output

import "hex/core/domain"

type Storage interface {
	Save(doc domain.Document) error
	FindByID(id int64) (*domain.Document, error)
}
