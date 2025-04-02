package domain

// DocumentRepo defines the interface for document storage
type DocumentRepo interface {
	Save(document Document) error
	FindByID(id int64) (*Document, error)
}
