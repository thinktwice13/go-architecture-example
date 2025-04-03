package document

import (
	"clean/domain/entity"
	"clean/domain/event"
	"clean/domain/repo"
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

func (uc *UploadUseCase) Upload(doc entity.Document) (int64, error) {
	if err := doc.ValidateForUpload(); err != nil {
		return 0, err // Return domain error directly
	}

	// Business logic ...

	_ = uc.repo.Save(doc)
	_ = uc.pub.Publish(event.DocumentUploaded{
		DocumentID: doc.ID,
		Timestamp:  time.Now(),
	})

	return doc.ID, nil
}
