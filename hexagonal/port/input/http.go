package input

import "hex/core/domain"

type Http interface {
	UploadDocument(doc domain.Document) error
}
