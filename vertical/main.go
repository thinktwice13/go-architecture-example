package main

import (
	"log"
	"net/http"
	"vert/feature/upload"
	"vert/shared/db"
)

func main() {
	store := db.InMemoryDB{}
	svc := upload.NewService(&store)
	handler := upload.NewHTTPHandler(svc)

	router := http.NewServeMux()
	router.HandleFunc("/upload", handler.HandleUpload)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
