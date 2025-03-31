package main

import (
	inadapter "hex/adapter/input"
	outadapter "hex/adapter/output"
	"hex/core/services"
	"log"
	"net/http"
)

func main() {
	store := &outadapter.DB{}
	svc := services.NewUploadService(store)
	handler := inadapter.NewDocumentHandler(svc)

	router := http.NewServeMux()
	inadapter.SetupRoutes(router, handler)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8084", router))
}
