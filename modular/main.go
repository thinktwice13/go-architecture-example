package main

import (
	"log"
	"modular/document"
	"modular/storage"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	db := storage.ConnectDB()
	router.Handle("/upload", document.NewUploadHandler(db))
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
