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

func (d *Document) HandleUpload(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_ = d.svc.Upload(model.Document{
		ID:        time.Now().UnixNano(),
		Name:      "sample.txt",
		Status:    "new",
		CreatedAt: time.Now().UTC(),
	})
	w.WriteHeader(http.StatusCreated)
}

func (d *Document) HandleFind(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	docID, _ := strconv.ParseInt(id, 10, 64)
	doc, _ := d.svc.FindByID(docID)
	_ = json.NewEncoder(w).Encode(doc)
}
