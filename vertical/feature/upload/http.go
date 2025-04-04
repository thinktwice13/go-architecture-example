// Package upload provides all the document uploading needs.
// This is an example of a more complex feature that has its own handler, service, and repo.
package upload

import (
	"net/http"
	"strconv"
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
		if err := r.ParseMultipartForm(2 << 20); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		_, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}

		if err := svc.upload(domain.Document{
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
	})
}
