package http

import (
	"clean/domain/entity"
	"encoding/json"
	"errors"
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
func NewDocumentHandler(service input.DocumentService, r *httprouter.Router) *DocumentHandler {
	h := &DocumentHandler{service}
	r.POST("/documents", h.handleUpload)
	r.GET("/documents/:id", h.handleGet)
	return h
}

func (h *DocumentHandler) handleUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	if err := h.service.Upload(domain.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      header.Filename,
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	}); err != nil {
		http.Error(w, "Error uploading the file", http.StatusInternalServerError)
	}

	w.Header().Set("Location", "/documents/"+header.Filename)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(header.Filename)))
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte(header.Filename)); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func (h *DocumentHandler) handleGet(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	doc, err := h.service.Find(ps.ByName("id"))
	switch {
	case errors.Is(err, entity.ErrNotFound):
		http.Error(w, "Document not found", http.StatusNotFound)
	case err != nil:
		http.Error(w, "Error retrieving the document", http.StatusInternalServerError)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(doc.Name)))
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(doc); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}
