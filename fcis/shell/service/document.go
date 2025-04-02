package service

import (
	"errors"
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
	validationResult := document.Validate(doc)
	if !validationResult.IsValid {
		return errors.New("validation failed")
	}

	// Storage depends on the infrastructure
	return ds.store.Save(doc)
}
