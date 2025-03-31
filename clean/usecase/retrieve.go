package usecase

import (
	"clean/entity"
	"clean/usecase/iface"
	"time"
)

type RetrieveDocumentUseCase struct {
	repo     iface.DocumentRepository
	eventBus iface.EventBus
}

func NewRetrieveDocumentUseCase(
	repo iface.DocumentRepository,
	eventBus iface.EventBus,
) *RetrieveDocumentUseCase {
	return &RetrieveDocumentUseCase{repo, eventBus}
}

func (uc *RetrieveDocumentUseCase) Retrieve(id int64) (*entity.Document, error) {
	doc, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	_ = uc.eventBus.Publish(entity.DocumentRetrieved{
		DocumentID: doc.ID,
		Timestamp:  time.Now(),
	})

	return doc, nil
}
