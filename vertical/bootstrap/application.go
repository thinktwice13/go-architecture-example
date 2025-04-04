package bootstrap

import (
	"net/http"
	"vertical/feature/retrieve"
	"vertical/feature/upload"
	"vertical/shared/config"
	"vertical/shared/db"
	"vertical/shared/event"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	_ = config.Load()
	store := db.Connect()

	// Common driven services and workers
	eb := &event.Bus{}

	// Core services with dependencies
	router := httprouter.New()
	upload.New(router, store, eb)
	retrieve.New(router, store)

	// Start
	return http.ListenAndServe(":8080", router)
}
