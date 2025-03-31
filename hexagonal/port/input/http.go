package input

import "hex/core/domain"

type Http interface {
	UploadDocument(doc domain.Document) error
	GetDocument(id int64) (*domain.Document, error)
}
