package api

// DocumentService defines the public API for the documents module
type DocumentService interface {
	UploadDocument(request UploadRequest) error
	GetDocument(id int64) (*DocumentDTO, error)
}
