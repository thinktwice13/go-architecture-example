package usecase

import (
	"clean/entity"
	"clean/usecase/iface"
)

type UploadDocumentUseCase struct {
	repo     iface.DocumentRepository
	eventBus iface.EventBus
}

func NewDocumentUploadUseCase(repo iface.DocumentRepository, eb iface.EventBus) *UploadDocumentUseCase {
	return &UploadDocumentUseCase{repo, eb}
}

func (uc *UploadDocumentUseCase) Upload(doc entity.Document) error {
	return uc.repo.Save(doc)
}
