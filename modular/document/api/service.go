package api

import "modular/document/domain"

// DocumentService defines the public API for the documents module
type DocumentService interface {
	UploadDocument(domain.Document) error
	GetDocument(id string) (*domain.Document, error)
}
