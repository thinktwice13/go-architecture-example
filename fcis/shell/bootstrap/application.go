package bootstrap

import (
	"fcis/shell/api"
	"fcis/shell/db"
	"fcis/shell/event"
	"fcis/shell/repo"
	"fcis/shell/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	dbConn := &db.Conn{}
	eb := &event.Bus{}

	// Domain
	docRepo := repo.NewDocumentRepo(dbConn)

	uploadService := service.NewUploadService(docRepo, eb)
	uploadHandler := api.NewUploadHandler(uploadService)

	retrieveService := service.NewRetrieveService(docRepo)
	retrieveHandler := api.NewRetrieveHandler(retrieveService)

	router := httprouter.New()
	router.POST("/documents", uploadHandler.Handle)
	router.GET("/documents/:id", retrieveHandler.Handle)
	return http.ListenAndServe(":8080", router)
}
