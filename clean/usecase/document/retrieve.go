package document

import (
	"clean/domain/entity"
)

// RetrieveUseCase implements the document retrieval use case
// Separation at use case level for distinct features
type RetrieveUseCase struct {
	repo Repo // Same repo interface, shared across use cases
	pub  EventPublisher
}

// NewRetrieveUseCase creates a new retrieve use case
func NewRetrieveUseCase(repo Repo) *RetrieveUseCase {
	return &RetrieveUseCase{
		repo: repo,
	}
}

// Retrieve fetches a document by ID
func (uc *RetrieveUseCase) Retrieve(id int64) (*entity.Document, error) {
	return uc.repo.FindByID(id)
}
