package document

import (
	"clean/domain/entity"
	"clean/domain/repo"
)

// RetrieveUseCase implements the document retrieval use case
// Separation at use case level for distinct features
type RetrieveUseCase struct {
	repo repo.DocRepo // Same repo interface, shared across use cases
}

func NewRetrieveUseCase(repo repo.DocRepo) *RetrieveUseCase {
	return &RetrieveUseCase{
		repo: repo,
	}
}

func (uc *RetrieveUseCase) Retrieve(id int64) (*entity.Document, error) {
	return uc.repo.FindByID(id)
}
