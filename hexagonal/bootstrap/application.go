package bootstrap

import (
	apphttp "hexagonal/adapter/input/http"
	"hexagonal/adapter/input/server"
	"hexagonal/adapter/output/db"
	"hexagonal/adapter/output/errs"
	events "hexagonal/adapter/output/event"
	"hexagonal/adapter/output/repo"
	"hexagonal/core/service"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	store := &db.Conn{}
	eb := &events.Bus{}

	// Services
	docRepo := repo.NewInMemory(store)
	errHandler := errs.NewHandler()
	svc := service.NewDocumentService(docRepo, eb, errHandler)
	httpHandler := apphttp.NewDocumentHandler(svc)
	router := httprouter.New()
	httpHandler.Routes(router)

	// Start
	srv := server.NewServer(router)
	return srv.Start()
}
