package persistence

//
// import (
// 	"clean/domain/entity"
// 	"clean/infra/storage"
// 	"testing"
// 	"time"
// )
//
// func TestDocumentRepositoryIntegration(t *testing.T) {
// 	// Routes real DB or init test container
// 	db := storage.ConnectDB()
// 	repo := NewDocumentRepository(db)
//
// 	// Create a test document
// 	doc := entity.Document{
// 		ID:        123,
// 		Name:      "integration.txt",
// 		Status:    "new",
// 		CreatedAt: time.Now(),
// 	}
//
// 	// Test Save
// 	err := repo.Save(doc)
// 	if err != nil {
// 		t.Fatalf("Failed to save document: %v", err)
// 	}
//
// 	// Test Retrieve
// 	retrieved, err := repo.FindByID(123)
// 	if err != nil {
// 		t.Fatalf("Failed to retrieve document: %v", err)
// 	}
//
// 	if retrieved.ID != doc.ID {
// 		t.Errorf("Expected ID %d, got %d", doc.ID, retrieved.ID)
// 	}
//
// 	if retrieved.Name != doc.Name {
// 		t.Errorf("Expected Name %s, got %s", doc.Name, retrieved.Name)
// 	}
// }
