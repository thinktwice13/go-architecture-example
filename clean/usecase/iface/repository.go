package iface

import "clean/entity"

type DocumentRepository interface {
	Save(document entity.Document) error
	FindByID(id int64) (*entity.Document, error)
}
