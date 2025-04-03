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
	store := &db.Conn{}
	docRepo := repo.NewDocumentRepo(store)
	svc := service.NewDocumentService(docRepo)
	httpHandler := handler.NewDocumentHandler(svc)
	router := httprouter.New()
	router.POST("/documents", httpHandler.HandleUpload)
	router.GET("/documents/:id", httpHandler.HandleFind)
	log.Fatal(http.ListenAndServe(":8080", router))
}
