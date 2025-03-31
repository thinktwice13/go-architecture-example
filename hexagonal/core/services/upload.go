package services

import (
	"hex/core/domain"
	"hex/port/input"
	"hex/port/output"
	"time"
)

var (
	_ input.Http = (*DocumentService)(nil)
)

type DocumentService struct {
	repo output.Storage
	publ output.EventPublisher
}

func NewDocumentService(repo output.Storage, eb output.EventPublisher) *DocumentService {
	return &DocumentService{repo, eb}
}

func (s *DocumentService) UploadDocument(doc domain.Document) error {
	_ = s.repo.Save(doc)
	s.publ.Publish(domain.DocumentUploaded{Document: doc, Timestamp: time.Now().UTC()})
	return nil
}

func (s *DocumentService) GetDocument(id int64) (*domain.Document, error) {
	doc, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
