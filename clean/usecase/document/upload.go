package document

import (
	"clean/domain/entity"
	"clean/domain/event"
	"time"
)

type UploadUseCase struct {
	repo Repo
	pub  EventPublisher
}

func NewUploadUseCase(repo Repo, events EventPublisher) *UploadUseCase {
	return &UploadUseCase{
		repo: repo,
		pub:  events,
	}
}

// UploadInput represents the input data for the upload use case
type UploadInput struct {
	Name string
}

func (uc *UploadUseCase) Upload(input UploadInput) (int64, error) {
	doc := entity.Document{
		ID:        time.Now().UnixNano(),
		Name:      input.Name,
		Status:    "new",
		CreatedAt: time.Now(),
	}

	// Validate the document
	if err := doc.ValidateForUpload(); err != nil {
		return 0, err // Return domain error directly
	}

	// Store the document
	_ = uc.repo.Save(doc)

	// Publish event
	_ = uc.pub.Publish(event.DocumentUploaded{
		DocumentID: doc.ID,
		Timestamp:  time.Now(),
	})

	return doc.ID, nil
}
