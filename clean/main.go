package main

import (
	"clean/adapter/api"
	"clean/adapter/eventbus"
	"clean/adapter/persistence"
	"clean/infra/storage"
	"clean/usecase"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	repo := persistence.NewDocumentRepository(storage.ConnectDB())
	eb := &eventbus.Bus{}
	upload := usecase.NewDocumentUploadUseCase(repo, eb)
	retrieval := usecase.NewRetrieveDocumentUseCase(repo, eb)
	apiHandler := api.NewHandler(upload, retrieval)

	router := httprouter.New()
	router.POST("/documents", apiHandler.UploadDocument)
	router.GET("/documents/:id", apiHandler.GetDocument)
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
