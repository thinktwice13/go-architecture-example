package api

import (
	"encoding/json"
	"fcis/shell/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// RetrieveHandler handles HTTP requests for document retrieval
type RetrieveHandler struct {
	retrieveService *service.RetrieveService
}

// NewRetrieveHandler creates a new retrieve handler
func NewRetrieveHandler(retrieveService *service.RetrieveService) *RetrieveHandler {
	return &RetrieveHandler{retrieveService: retrieveService}
}

// Handle processes the HTTP request
func (h *RetrieveHandler) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, _ := strconv.ParseInt(id, 10, 64)
	doc, _ := h.retrieveService.GetDocument(docID)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(doc)
}
