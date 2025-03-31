package upload

import (
	"hex/core/domain"
	"vert/shared/db"
)

type Service struct {
	repo *Repository
}

func NewService(db *db.InMemoryDB) *Service {
	return &Service{NewRepository(db)}
}

func (s *Service) UploadDocument(doc domain.Document) error {
	return s.repo.Save(doc)
}
