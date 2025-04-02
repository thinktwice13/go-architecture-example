package document

import (
	"clean/domain/entity"
	"clean/domain/event"
	"clean/domain/repo"
	"time"
)

type UploadUseCase struct {
	repo repo.DocRepo
	pub  event.Publisher
}

func NewUploadUseCase(repo repo.DocRepo, pub event.Publisher) *UploadUseCase {
	return &UploadUseCase{repo, pub}
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

	if err := doc.ValidateForUpload(); err != nil {
		return 0, err // Return domain error directly
	}

	_ = uc.repo.Save(doc)
	_ = uc.pub.Publish(event.DocumentUploaded{
		DocumentID: doc.ID,
		Timestamp:  time.Now(),
	})

	return doc.ID, nil
}
