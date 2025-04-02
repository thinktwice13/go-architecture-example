// Package upload provides all the document uploading needs.
// This is an example of a more complex feature that has its own handler, service, and repo.
package upload

import (
	"net/http"
	"time"
	"vert/shared/db"
	"vert/shared/domain"
	"vert/shared/event"

	"github.com/julienschmidt/httprouter"
)

// New bootstraps the upload feature from low-level infrastructure components, including the router attachment.
func New(router *httprouter.Router, db *db.Conn, pub *event.Bus) {
	repo := &Repo{db}
	svc := Service{repo, pub}

	router.POST("/documents", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Validation...
		doc := domain.Document{
			ID:        time.Now().UnixNano(),
			Name:      "sample.txt",
			Status:    "new",
			CreatedAt: time.Now().UTC(),
		}

		_ = svc.UploadDocument(doc)
		_ = pub.Publish(domain.DocumentEvent{
			Type:      domain.EventDocumentUploaded,
			Document:  &doc,
			Timestamp: time.Time{}.UTC(),
		})

		w.WriteHeader(http.StatusCreated)
	})
}
