package bootstrap

import (
	"hexagonal/adapter/eventbus"
	apphttp "hexagonal/adapter/input/http"
	"hexagonal/adapter/input/server"
	"hexagonal/adapter/output/db"
	"hexagonal/adapter/output/errs"
	"hexagonal/adapter/output/repo"
	"hexagonal/core/config"
	"hexagonal/core/service"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	_ = config.Load()
	store := db.Connect()

	// Driven Services
	eb := &eventbus.Bus{}
	docRepo := repo.NewRepo(store)
	errHandler := errs.NewHandler()

	// Core Services
	svc := service.NewDocumentService(docRepo, eb, errHandler)

	// Drivers
	router := httprouter.New()
	apphttp.NewDocumentHandler(svc, router)

	// Start
	return server.NewServer(router).Start()
}
