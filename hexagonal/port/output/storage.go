package output

import "hex/core/domain"

type Storage interface {
	Save(doc domain.Document) error
}
