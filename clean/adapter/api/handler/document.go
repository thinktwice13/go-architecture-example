package handler

import (
	"clean/domain/entity"
	"clean/usecase/document"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// DocumentHandler handles HTTP requests related to documents
// Each handler is responsible for a specific domain area, but may use multiple use cases
type DocumentHandler struct {
	uploads    *document.UploadUseCase
	retrievals *document.RetrieveUseCase
}

func NewDocumentHandler(
	upload *document.UploadUseCase,
	retrieval *document.RetrieveUseCase,
) *DocumentHandler {
	return &DocumentHandler{upload, retrieval}
}

func (h *DocumentHandler) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	if err := h.uploads.Upload(entity.Document{
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

func (h *DocumentHandler) HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	doc, err := h.retrievals.Find(ps.ByName("id"))
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
