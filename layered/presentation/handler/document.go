package handler

import (
	"encoding/json"
	"layered/business/service"
	"layered/data/model"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Document struct {
	svc *service.DocumentService
}

func NewDocumentHandler(svc *service.DocumentService) *Document {
	return &Document{svc: svc}
}

func (d *Document) HandleUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	if err := d.svc.Upload(model.Document{
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

func (d *Document) HandleFind(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	doc, _ := d.svc.FindByID(ps.ByName("id"))
	_ = json.NewEncoder(w).Encode(doc)
}
