package api

import (
	"encoding/json"
	"modular/document/domain"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// HTTPHandler handles HTTP requests for documents
type HTTPHandler struct {
	documentService DocumentService
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler(documentService DocumentService) *HTTPHandler {
	return &HTTPHandler{
		documentService: documentService,
	}
}

// RegisterRoutes registers HTTP routes
func (h *HTTPHandler) RegisterRoutes(router *httprouter.Router) {
	router.POST("/documents", h.uploadDocument)
	router.GET("/documents/:id", h.getDocument)
}

// uploadDocument handles document upload requests
func (h *HTTPHandler) uploadDocument(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = h.documentService.UploadDocument(domain.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	})
	w.WriteHeader(http.StatusCreated)
}

// getDocument handles document retrieval requests
func (h *HTTPHandler) getDocument(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	document, _ := h.documentService.GetDocument(ps.ByName("id"))
	_ = json.NewEncoder(w).Encode(document)
}
