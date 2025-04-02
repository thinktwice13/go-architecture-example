package api

import "time"

// DocumentDTO is a data transfer object for documents
type DocumentDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// UploadRequest represents a document upload request
type UploadRequest struct {
	Name string `json:"name"`
}

// UploadResponse represents a document upload response
type UploadResponse struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
