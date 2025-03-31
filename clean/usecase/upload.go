package usecase

import (
	"clean/entity"
	"clean/usecase/iface"
)

type UploadDocumentUseCase struct {
	repo iface.DocumentRepository
}

func NewDocumentUploadUseCase(repo iface.DocumentRepository) *UploadDocumentUseCase {
	return &UploadDocumentUseCase{repo}
}

func (uc *UploadDocumentUseCase) Upload(doc entity.Document) error {
	return uc.repo.Save(doc)
}
