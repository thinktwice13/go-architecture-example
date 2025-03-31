package input

import (
	"encoding/json"
	"hex/core/domain"
	"hex/port/input"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type DocumentHandler struct {
	documentService input.Http
}

func NewDocumentHandler(service input.Http) *DocumentHandler {
	return &DocumentHandler{
		documentService: service,
	}
}

func (h *DocumentHandler) HandleUploadDocument(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	doc := domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		CreatedAt: time.Now(),
	}

	if err := h.documentService.UploadDocument(doc); err != nil {
		http.Error(w, "Upload failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *DocumentHandler) HandleGetDocument(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid document ID", http.StatusBadRequest)
		return
	}

	doc, err := h.documentService.GetDocument(docID)
	if err != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(doc)
}

func SetupRoutes(router *httprouter.Router, documentHandler *DocumentHandler) {
	router.POST("/documents", documentHandler.HandleUploadDocument)
	router.GET("/documents/:id", documentHandler.HandleGetDocument)
}
