package main

import (
	"log"
	"net/http"
	"vert/feature/retrieve"
	"vert/feature/upload"
	"vert/shared/db"
	"vert/shared/event"

	"github.com/julienschmidt/httprouter"
)

func main() {
	store := &db.DB{}
	eb := &event.Bus{}
	router := httprouter.New()
	router.POST("/documents", upload.New(store, eb))
	router.GET("/documents/:id", retrieve.New(store))
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
