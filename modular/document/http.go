package document

import (
	"modular/storage"
	"net/http"
	"time"
)

func NewUploadHandler(db *storage.DB) http.HandlerFunc {
	repo := InMemoryRepository{db}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		doc := Document{
			ID:        time.Now().UnixNano(),
			Name:      r.FormValue("sample.txt"),
			Status:    "new",
			CreatedAt: time.Now(),
		}

		if err := repo.Save(doc); err != nil {
			http.Error(w, "Failed to save document", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
