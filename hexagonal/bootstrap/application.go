package bootstrap

import (
	apphttp "hex/adapter/input/http"
	"hex/adapter/output/db"
	"hex/adapter/output/errs"
	events "hex/adapter/output/event"
	"hex/adapter/output/repo"
	"hex/core/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	store := &db.Conn{}
	eb := events.Bus{}

	// Services
	repo := repo.NewInMemoryRepo(store)
	errHandler := errs.NewErrorHandler()
	svc := service.NewDocumentService(repo, &eb, errHandler)
	httpHandler := apphttp.NewDocumentHandler(svc)
	router := httprouter.New()
	httpHandler.Routes(router)

	// Start
	return http.ListenAndServe(":8080", router)
}
