package upload

import (
	"hex/core/domain"
	"net/http"
	"time"
)

type HTTPHandler struct {
	svc *Service
}

func NewHTTPHandler(svc *Service) *HTTPHandler {
	return &HTTPHandler{svc}
}

func (h *HTTPHandler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	doc := domain.Document{
		ID:        time.Now().UnixNano(),
		Name:      "example.txt",
		Status:    "new",
		CreatedAt: time.Now(),
	}
	if err := h.svc.UploadDocument(doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
