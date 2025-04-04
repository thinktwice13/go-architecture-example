package upload

import (
	"encoding/json"
	"time"
	"vertical/shared/domain"
	"vertical/shared/event"
)

type Service struct {
	repo *Repo
	pub  *event.Bus
}

func (s *Service) upload(doc domain.Document) error {
	if err := s.repo.Save(doc); err != nil {
		return err
	}

	if bytes, err := json.Marshal(doc); err != nil {
		return err
	} else {
		return s.pub.Publish(domain.DocumentEvent{
			Type:      domain.EventDocumentUploaded,
			Timestamp: time.Now().UTC(),
			Data:      bytes,
		})
	}
}
