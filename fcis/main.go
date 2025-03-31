package main

import (
	"fcis/shell/api"
	"fcis/shell/infra"
	"fcis/shell/service"
	"fcis/shell/storage"
	"log"
	"net/http"
)

func main() {
	db := infra.ConnectDB()
	store := storage.NewDocumentStore(db)
	svc := service.NewDocumentService(store)
	handler := api.NewDocumentHandler(svc)

	router := http.NewServeMux()
	router.HandleFunc("/upload", handler.HandleUpload)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8083", router))
}
