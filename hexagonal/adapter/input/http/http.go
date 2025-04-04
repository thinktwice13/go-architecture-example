package http

import (
	"encoding/json"
	"hexagonal/core/domain"
	"hexagonal/port/input"
	"net/http"
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
	r.POST("/documents", h.upload)
	r.GET("/documents/:id", h.find)
	return h
}

func (h *DocumentHandler) upload(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = h.service.Upload(domain.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      "sample.txt",
		CreatedAt: time.Now().UTC(),
	})

	w.WriteHeader(http.StatusCreated)
}

func (h *DocumentHandler) find(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	doc, _ := h.service.Find(ps.ByName("id"))
	_ = json.NewEncoder(w).Encode(doc)
}
