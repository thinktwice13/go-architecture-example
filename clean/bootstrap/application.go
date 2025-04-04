package bootstrap

import (
	"clean/adapter/api/handler"
	"clean/adapter/event"
	"clean/adapter/repo"
	"clean/infra/config"
	"clean/infra/db"
	"clean/infra/server"
	"clean/usecase/document"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	_ = config.Load()
	dbConn := db.Connect()

	// Storage and outputs
	eb := &event.Bus{}
	docRepo := repo.NewDocumentRepo(dbConn)

	// Usecases
	uploadUC := document.NewUploadUseCase(docRepo, eb)
	retrieveUC := document.NewRetrieveUseCase(docRepo)

	// Inputs
	router := httprouter.New()
	httpHandler := handler.NewDocumentHandler(uploadUC, retrieveUC)
	router.POST("/documents", httpHandler.Upload)
	router.GET("/documents/:id", httpHandler.HandleGet)

	// Server
	srv := server.NewServer(router)
	return srv.Start()
}
