package upload

import (
	"hex/core/domain"
	"time"
	"vert/shared/db"
	"vert/shared/event"
)

type Service struct {
	repo *Repository
	pub  event.Publisher
}

func newService(db *db.DB, pub event.Publisher) *Service {
	return &Service{newRepo(db), pub}
}

func (s *Service) UploadDocument(doc domain.Document) error {
	_ = s.repo.Save(doc)
	_ = s.pub.Publish(domain.DocumentUploaded{
		Document:  doc,
		Timestamp: time.Now().UTC(),
	})
	return nil
}
