package handler

import (
	"encoding/json"
	"layered/business/service"
	"layered/data/model"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Document struct {
	svc *service.DocumentService
}

func NewDocumentHandler(svc *service.DocumentService) *Document {
	return &Document{svc: svc}
}

func (d *Document) HandleUpload(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = d.svc.Upload(model.Document{
		ID:        time.Now().Format("20060102150405"),
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	})
	w.WriteHeader(http.StatusCreated)
}

func (d *Document) HandleFind(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	doc, _ := d.svc.FindByID(ps.ByName("id"))
	_ = json.NewEncoder(w).Encode(doc)
}
