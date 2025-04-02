package upload

import (
	"time"
	"vertical/shared/domain"
	"vertical/shared/event"
)

type Service struct {
	repo *Repo
	pub  *event.Bus
}

func (s *Service) UploadDocument(doc domain.Document) error {
	_ = s.repo.Save(doc)
	_ = s.pub.Publish(domain.DocumentEvent{
		Type:      domain.EventDocumentUploaded,
		Document:  &doc,
		Timestamp: time.Time{},
	})
	return nil
}
