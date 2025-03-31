package main

import (
	"clean/adapter/api"
	"clean/adapter/persistence"
	"clean/infra/storage"
	"clean/usecase"
	"log"
	"net/http"
)

func main() {
	repo := persistence.NewDocumentRepository(storage.ConnectDB())
	uc := usecase.NewDocumentUploadUseCase(repo)
	documentHandler := api.NewHandler(uc)

	router := http.NewServeMux()
	router.HandleFunc("/", documentHandler.UploadDocument)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8081", router))
}
