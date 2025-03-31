package api

import (
	"clean/entity"
	"clean/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type HTTPHandler struct {
	upload    *usecase.UploadDocumentUseCase
	retrieval *usecase.RetrieveDocumentUseCase
}

func NewHandler(
	upload *usecase.UploadDocumentUseCase,
	retrieval *usecase.RetrieveDocumentUseCase,
) *HTTPHandler {
	return &HTTPHandler{upload, retrieval}
}

func (h *HTTPHandler) UploadDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request to get the document
	// Call the use case to upload the document
	doc := entity.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		CreatedAt: time.Now(),
		Status:    "new",
	}

	if err := h.upload.Upload(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) GetDocument(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid document ID", http.StatusBadRequest)
		return
	}

	doc, err := h.retrieval.Retrieve(docID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrDocumentNotFound):
			http.Error(w, "Document not found", http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(doc)
}
