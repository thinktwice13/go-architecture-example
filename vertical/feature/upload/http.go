package upload

import (
	"hex/core/domain"
	"net/http"
	"time"
	"vert/shared/db"
	"vert/shared/event"

	"github.com/julienschmidt/httprouter"
)

func New(store *db.DB, pub event.Publisher) httprouter.Handle {
	svc := newService(store, pub)
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		doc := domain.Document{
			ID:        time.Now().UnixNano(),
			Name:      "example.txt",
			Status:    "new",
			CreatedAt: time.Now(),
		}

		if err := svc.UploadDocument(doc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
