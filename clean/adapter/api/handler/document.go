package handler

import (
	"clean/domain/entity"
	"clean/usecase/document"
	"encoding/json"
	"net/http"
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
	_, _ = h.uploads.Upload(entity.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	})
	w.WriteHeader(http.StatusCreated)
}

func (h *DocumentHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	doc, _ := h.retrievals.Retrieve(ps.ByName("id"))
	_ = json.NewEncoder(w).Encode(doc)
}
