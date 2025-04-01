package retrieve

import (
	"encoding/json"
	"net/http"
	"vert/shared/db"

	"github.com/julienschmidt/httprouter"
)

func New(db *db.DB) httprouter.Handle {
	// Vertical slice intentionally skipping service layer for simple features, without extensive validation or orchestration logic
	repo := &Repository{db}
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")
		doc, _ := repo.FindByID(id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(doc)
	}
}
