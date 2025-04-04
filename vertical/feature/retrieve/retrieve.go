// Package retrieve retrieves a single document by ID
// Vertical slices adapt to each feature's needs. Find is an example of a simpler feature that doesn't require a service layer and uses different pattern to bootstrap itself.
package retrieve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"vertical/shared/db"
	"vertical/shared/domain"

	"github.com/julienschmidt/httprouter"
)

type repo struct {
	db *db.Conn
}

func (r *repo) findByID(id string) (*domain.Document, error) {
	fmt.Println("Getting document from database")
	return &domain.Document{
		ID:        id,
		Name:      "sample.txt",
		Status:    "done",
		CreatedAt: time.Now().UTC(),
	}, nil
}

// New bootstraps the upload feature from low-level infrastructure components, including the router attachment.
func New(router *httprouter.Router, db *db.Conn) {
	docrepo := &repo{db}

	router.GET("/documents/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		doc, _ := docrepo.findByID(ps.ByName("id"))
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(doc)
	})
}
