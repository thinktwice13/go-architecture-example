package api

import (
	"clean/entity"
	"clean/usecase"
	"net/http"
	"time"
)

type HTTPHandler struct {
	upload *usecase.UploadDocumentUseCase
}

func NewHandler(upload *usecase.UploadDocumentUseCase) *HTTPHandler { return &HTTPHandler{upload} }

func (h *HTTPHandler) UploadDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request to get the document
	// Call the use case to upload the document
	doc := entity.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		CreatedAt: time.Now(),
		Status:    "new",
	}

	if err := h.upload.Upload(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
