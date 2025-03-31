package main

import (
	inadapter "hex/adapter/input"
	outadapter "hex/adapter/output"
	"hex/core/services"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	store := &outadapter.DB{}
	eb := &outadapter.EventBus{}
	svc := services.NewDocumentService(store, eb)
	handler := inadapter.NewDocumentHandler(svc)

	router := httprouter.New()
	inadapter.SetupRoutes(router, handler)
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
