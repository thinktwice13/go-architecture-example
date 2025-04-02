package http

import (
	"encoding/json"
	"hexagonal/core/domain"
	"hexagonal/port/input"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// DocumentHandler adapts HTTP requests to the document service input port
// This is a primary adapter that drives the application
type DocumentHandler struct {
	service input.DocumentService
}

// NewDocumentHandler creates a new document HTTP handler
func NewDocumentHandler(service input.DocumentService) *DocumentHandler {
	return &DocumentHandler{service}
}

// Routes registers the HTTP routes for document operations
func (h *DocumentHandler) Routes(r *httprouter.Router) {
	r.POST("/documents", h.upload)
	r.GET("/documents/:id", h.find)
}

// upload handles HTTP requests for document upload
func (h *DocumentHandler) upload(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = h.service.Upload(domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "sample.txt",
		CreatedAt: time.Now().UTC(),
	})

	w.WriteHeader(http.StatusCreated)
}

// find handles HTTP requests for document retrieval
func (h *DocumentHandler) find(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, _ := strconv.ParseInt(id, 10, 64)
	doc, _ := h.service.Find(docID)
	_ = json.NewEncoder(w).Encode(doc)
}
