package usecase

import (
	"clean/entity"
	"errors"
	"testing"
	"time"
)

// MockRepo is a mock implementation of DocumentRepository
type MockRepo struct {
	docs map[int64]entity.Document
}

func (m *MockRepo) Save(doc entity.Document) error {
	m.docs[doc.ID] = doc
	return nil
}

func (m *MockRepo) FindByID(id int64) (*entity.Document, error) {
	doc, exists := m.docs[id]
	if !exists {
		return nil, entity.ErrDocumentNotFound
	}
	return &doc, nil
}

// MockEventBus is a mock implementation of EventBus
type MockEventBus struct {
	events []entity.DocumentEvent
}

func (m *MockEventBus) Publish(event entity.DocumentEvent) error {
	m.events = append(m.events, event)
	return nil
}

func TestRetrieveDocument(t *testing.T) {
	// Setup
	mockRepo := &MockRepo{
		docs: map[int64]entity.Document{
			1: {
				ID:        1,
				Name:      "test.txt",
				Status:    "new",
				CreatedAt: time.Now(),
			},
		},
	}
	mockEventBus := &MockEventBus{}

	useCase := NewRetrieveDocumentUseCase(mockRepo, mockEventBus)

	// Test successful retrieval
	t.Run("Document exists", func(t *testing.T) {
		doc, err := useCase.Retrieve(1)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if doc == nil {
			t.Fatal("Expected document, got nil")
		}

		if doc.ID != 1 {
			t.Errorf("Expected ID 1, got %d", doc.ID)
		}

		if len(mockEventBus.events) != 1 {
			t.Errorf("Expected 1 event, got %d", len(mockEventBus.events))
		}

		event, ok := mockEventBus.events[0].(entity.DocumentRetrieved)
		if !ok {
			t.Error("Expected DocumentRetrieved event")
		}

		if event.DocumentID != 1 {
			t.Errorf("Expected event document ID 1, got %d", event.DocumentID)
		}
	})

	t.Run("Document not found", func(t *testing.T) {
		_, err := useCase.Retrieve(999)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, entity.ErrDocumentNotFound) {
			t.Errorf("Expected ErrDocumentNotFound, got %v", err)
		}
	})
}
