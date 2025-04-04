package api

import "modular/document/domain"

// DocumentService defines the public API for the documents module
type DocumentService interface {
	Upload(domain.Document) error
	Find(id string) (*domain.Document, error)
}
