package document

import (
	"clean/domain/entity"
	"clean/domain/event"
	"clean/domain/repo"
	"encoding/json"
	"time"
)

// UploadUseCase handles the business logic for uploading documents
// Separation at use case level for distinct features
type UploadUseCase struct {
	repo repo.DocRepo
	pub  event.Publisher
}

func NewUploadUseCase(repo repo.DocRepo, pub event.Publisher) *UploadUseCase {
	return &UploadUseCase{repo, pub}
}

func (uc *UploadUseCase) Upload(doc entity.Document) error {
	if err := doc.ValidateForUpload(); err != nil {
		return err // Return domain error directly
	}

	// Business logic ...

	if err := uc.repo.Save(doc); err != nil {
		return err // Return domain error directly
	}

	bytes, _ := json.Marshal(doc)
	return uc.pub.Publish(event.DocumentEvent{
		Type:      event.EventDocumentUploaded,
		Timestamp: time.Now(),
		Data:      bytes,
	})
}
