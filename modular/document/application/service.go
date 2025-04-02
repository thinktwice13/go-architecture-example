package application

import (
	commondomain "modular/common/domain"
	"modular/document/api"
	"modular/document/domain"
	"time"
)

// UploadService implements document upload functionality
type UploadService struct {
	documentRepo domain.DocumentRepo
	bus          EventPublisher
}

// EventPublisher defines the event publishing interface
type EventPublisher interface {
	Publish(event commondomain.DomainEvent) error
}

var _ api.DocumentService = (*UploadService)(nil)

// NewUploadService creates a new upload service
func NewUploadService(repo domain.DocumentRepo, eventBus EventPublisher) *UploadService {
	return &UploadService{
		documentRepo: repo,
		bus:          eventBus,
	}
}

// UploadDocument handles the document upload process
func (s *UploadService) UploadDocument(request api.UploadRequest) error {
	doc := domain.NewDocument(request.Name)
	doc.ID = time.Now().UnixNano()

	// Save the document to the repository
	_ = s.documentRepo.Save(*doc)

	// Publish an event after saving the document
	evt := domain.DocumentUploaded{
		BaseEvent: commondomain.BaseEvent{
			Type:      "DocumentUploaded",
			Timestamp: time.Now().UTC(),
		},
		DocumentID: doc.ID,
	}

	_ = s.bus.Publish(evt)
	return nil
}

// GetDocument retrieves a document by its ID
func (s *UploadService) GetDocument(id int64) (*api.DocumentDTO, error) {
	doc, err := s.documentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &api.DocumentDTO{
		ID:        doc.ID,
		Name:      doc.Name,
		Status:    doc.Status,
		CreatedAt: doc.CreatedAt,
	}, nil
}
