// Package retrieve retrieves a single document by ID
// Vertical slices adapt to each feature's needs. Retrieve is an example of a simple feature that doesn't require a service layer.
// It also initializes the feature in a different pattern, attaches the handler and returns nothing
package retrieve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"vert/shared/db"
	"vert/shared/domain"

	"github.com/julienschmidt/httprouter"
)

type Repo struct {
	db *db.Conn
}

func (r *Repo) FindByID(id int64) (*domain.Document, error) {
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
	repo := &Repo{db}

	router.GET("/documents/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		docID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		doc, _ := repo.FindByID(docID)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(doc)
	})
}
