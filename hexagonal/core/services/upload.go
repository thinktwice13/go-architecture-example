package services

import (
	"hex/core/domain"
	"hex/ports/input"
	"hex/ports/output"
)

var _ input.Http = (*UploadService)(nil)

type UploadService struct {
	repo output.Storage
}

func NewUploadService(repo output.Storage) *UploadService {
	return &UploadService{repo}
}

func (s *UploadService) UploadDocument(doc domain.Document) error {
	return s.repo.Save(doc)
}
