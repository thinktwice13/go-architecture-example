package api

import (
	"fcis/core/document"
	"fcis/shell/service"
	"net/http"
	"time"
)

type DocumentHandler struct {
	service *service.DocumentService
}

func NewDocumentHandler(service *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		service: service,
	}
}

func (h *DocumentHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Call service (side effect)
	doc := document.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}

	if err := h.service.UploadDocument(doc); err != nil {
		http.Error(w, "Failed to upload document", http.StatusInternalServerError)
		return
	}

	// Response is a side effect
	w.WriteHeader(http.StatusCreated)
}
