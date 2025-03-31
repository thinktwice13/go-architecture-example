package input

import (
	"hex/core/domain"
	"hex/ports/input"
	"net/http"
	"time"
)

type DocumentHandler struct {
	documentService input.Http
}

func NewDocumentHandler(service input.Http) *DocumentHandler {
	return &DocumentHandler{
		documentService: service,
	}
}

func (h *DocumentHandler) HandleUploadDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	doc := domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		CreatedAt: time.Now(),
	}

	if err := h.documentService.UploadDocument(doc); err != nil {
		http.Error(w, "Upload failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SetupRoutes(router *http.ServeMux, documentHandler *DocumentHandler) {
	router.HandleFunc("/documents", documentHandler.HandleUploadDocument)
}
