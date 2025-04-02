package api

import (
	"fcis/core/document"
	"fcis/shell/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UploadHandler handles HTTP requests for document upload
// Part of the Imperative Shell - manages side effects
type UploadHandler struct {
	uploadService *service.UploadService
}

// NewUploadHandler creates a new upload handler
func NewUploadHandler(uploadService *service.UploadService) *UploadHandler {
	return &UploadHandler{uploadService: uploadService}
}

// Handle processes the HTTP request
func (h *UploadHandler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = h.uploadService.UploadDocument(document.New("sample.txt"))

	w.WriteHeader(http.StatusCreated)
}
