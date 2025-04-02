package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	_ = h.documentService.UploadDocument(UploadRequest{
		Name: "sample.txt",
	})
	w.WriteHeader(http.StatusCreated)
}

// getDocument handles document retrieval requests
func (h *HTTPHandler) getDocument(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, _ := strconv.ParseInt(id, 10, 64)
	document, _ := h.documentService.GetDocument(docID)
	_ = json.NewEncoder(w).Encode(document)
}
