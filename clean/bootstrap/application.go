package bootstrap

import (
	"clean/adapter/api/handler"
	"clean/adapter/event"
	"clean/adapter/persistence"
	"clean/infra/db"
	"clean/infra/server"
	"clean/usecase/document"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	dbConn := &db.Conn{}
	eb := &event.Bus{}

	// Domain
	// Repo
	docRepo := persistence.NewDocumentRepo(dbConn)

	// Usecases
	uploadUC := document.NewUploadUseCase(docRepo, eb)
	retrieveUC := document.NewRetrieveUseCase(docRepo)

	// Http
	router := httprouter.New()
	httpHandler := handler.NewDocumentHandler(uploadUC, retrieveUC)
	router.POST("/documents", httpHandler.Upload)
	router.GET("/documents/:id", httpHandler.Get)

	// Server
	srv := server.NewServer(router)
	return srv.Start()
}
