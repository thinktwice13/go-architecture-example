package application

import (
	commondomain "modular/common/domain"
	"modular/document/api"
	"modular/document/domain"
	"time"
)

// DocumentService implements document upload functionality
type DocumentService struct {
	documentRepo domain.DocumentRepo
	bus          EventPublisher
}

// EventPublisher defines the event publishing interface
type EventPublisher interface {
	Publish(event commondomain.Event) error
}

var _ api.DocumentService = (*DocumentService)(nil)

// NewDocumentService creates a new upload service
func NewDocumentService(repo domain.DocumentRepo, eventBus EventPublisher) *DocumentService {
	return &DocumentService{
		documentRepo: repo,
		bus:          eventBus,
	}
}

// UploadDocument handles the document upload process
func (s *DocumentService) UploadDocument(request api.UploadRequest) error {
	doc := domain.NewDocument(request.Name)
	doc.ID = time.Now().UnixNano()

	// Save the document to the repository
	_ = s.documentRepo.Save(*doc)

	// Publish an event after saving the document
	evt := domain.DocumentUploaded{
		BaseEvent: commondomain.BaseEvent{
			Type:      commondomain.EventDocumentUploaded,
			Timestamp: time.Now().UTC(),
		},
		DocumentID: doc.ID,
	}

	_ = s.bus.Publish(evt)
	return nil
}

// GetDocument retrieves a document by its ID
func (s *DocumentService) GetDocument(id int64) (*api.DocumentDTO, error) {
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
