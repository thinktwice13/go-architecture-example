package service

import (
	"fcis/core/document"
)

type DocumentStore interface {
	Save(doc document.Document) error
}

type DocumentService struct {
	store DocumentStore
}

func NewDocumentService(store DocumentStore) *DocumentService {
	return &DocumentService{store}
}

func (ds *DocumentService) UploadDocument(doc document.Document) error {
	// Validate is a core function without side effects
	if err := document.Validate(doc); err != nil {
		return err
	}

	// Storage depends on the infrastructure
	return ds.store.Save(doc)
}
