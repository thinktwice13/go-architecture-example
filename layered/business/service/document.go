package service

import (
	"layered/data/model"
	"layered/data/repo"
)

type DocumentService struct {
	repo *repo.Document
}

func NewDocumentService(repo *repo.Document) *DocumentService {
	return &DocumentService{repo}
}

func (s *DocumentService) Upload(doc model.Document) error {
	// Validate
	if doc.Name == "" {
		return ErrInvalidDocument{Field: "name", Message: "Document name cannot be empty"}
	}

	return s.repo.Save(doc)
}

func (s *DocumentService) FindByID(id int64) (*model.Document, error) {
	return s.repo.FindByID(id)
}

type ErrInvalidDocument struct {
	Field   string
	Message string
}

func (e ErrInvalidDocument) Error() string {
	return e.Message
}
