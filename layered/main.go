package main

import (
	"layered/business/service"
	"layered/data/db"
	"layered/data/repo"
	"layered/presentation/handler"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Infra
	store := db.Connect()
	// No eventbus in this example

	// Data
	docRepo := repo.NewDocumentRepo(store)

	// Business
	svc := service.NewDocumentService(docRepo)

	// Presentation
	httpHandler := handler.NewDocumentHandler(svc)
	router := httprouter.New()
	router.POST("/documents", httpHandler.HandleUpload)
	router.GET("/documents/:id", httpHandler.HandleFind)

	// Start
	log.Fatal(http.ListenAndServe(":8080", router))
}
