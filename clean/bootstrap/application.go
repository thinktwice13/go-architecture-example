package bootstrap

import (
	"clean/adapter/api/handler"
	"clean/adapter/event"
	"clean/adapter/persistence"
	"clean/infra/db"
	"clean/usecase/document"
	"net/http"

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
	return http.ListenAndServe(":8080", router)
}
