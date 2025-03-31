package retrieve

import (
	"encoding/json"
	"net/http"
	"vert/shared/db"

	"github.com/julienschmidt/httprouter"
)

func New(db *db.InMemoryDB) httprouter.Handle {
	// Intentionally skipping service layer for simple featuress without extensive validation or orchestration needed
	repo := &Repository{db}
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		doc, err := repo.FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(doc)
	}
}
