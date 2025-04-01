package handler

import (
	"clean/usecase/document"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DocumentHandler handles HTTP requests related to documents
// Each handler is responsible for a specific domain area, but may use multiple use cases
type DocumentHandler struct {
	uploadUseCase   *document.UploadUseCase
	retrieveUseCase *document.RetrieveUseCase
}

func NewDocumentHandler(
	uploadUseCase *document.UploadUseCase,
	retrieveUseCase *document.RetrieveUseCase,
) *DocumentHandler {
	return &DocumentHandler{
		uploadUseCase:   uploadUseCase,
		retrieveUseCase: retrieveUseCase,
	}
}

func (h *DocumentHandler) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = h.uploadUseCase.Upload(document.UploadInput{Name: "sample.txt"})
	w.WriteHeader(http.StatusCreated)
}

func (h *DocumentHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, _ := strconv.ParseInt(id, 10, 64)
	doc, _ := h.retrieveUseCase.Retrieve(r.Context(), docID)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(doc)
}
