// Package upload provides all the document uploading needs.
// This is an example of a more complex feature that has its own handler, service, and repo.
package upload

import (
	"net/http"
	"time"
	"vertical/shared/db"
	"vertical/shared/domain"
	"vertical/shared/event"

	"github.com/julienschmidt/httprouter"
)

// New bootstraps the upload feature from low-level infrastructure components, including the router attachment.
func New(router *httprouter.Router, db *db.Conn, pub *event.Bus) {
	repo := &Repo{db}
	svc := Service{repo, pub}

	router.POST("/documents", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Validation...
		doc := domain.Document{
			ID:        time.Now().Format("20060102150405"),
			Name:      "sample.txt",
			Status:    "new",
			CreatedAt: time.Now().UTC(),
		}

		_ = svc.UploadDocument(doc)
		w.WriteHeader(http.StatusCreated)
	})
}
