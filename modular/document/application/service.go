package application

import (
	"encoding/json"
	commondomain "modular/common/domain"
	"modular/document/api"
	"modular/document/domain"
	"time"
)

type DocumentService struct {
	documentRepo domain.DocumentRepo
	bus          EventPublisher
}

type EventPublisher interface {
	Publish(event commondomain.Event) error
}

var _ api.DocumentService = (*DocumentService)(nil)

func NewDocumentService(repo domain.DocumentRepo, eventBus EventPublisher) *DocumentService {
	return &DocumentService{repo, eventBus}
}

func (s *DocumentService) Upload(doc domain.Document) error {
	if err := s.documentRepo.Save(doc); err != nil {
		return err
	}

	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	// Uses the common event struct, but local domain event type
	if err := s.bus.Publish(commondomain.Event{
		EventType: domain.EventDocumentUploaded,
		Timestamp: time.Now().UTC(),
		Data:      data,
	}); err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) Find(id string) (*domain.Document, error) {
	doc, _ := s.documentRepo.FindByID(id)
	return doc, nil
}
