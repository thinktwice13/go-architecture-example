package application

import (
	"net/http"
	"vert/feature/retrieve"
	"vert/feature/upload"
	"vert/shared/db"
	"vert/shared/event"

	"github.com/julienschmidt/httprouter"
)

func BootstrapAndRun() error {
	// Infra
	router := httprouter.New() // input
	store := &db.Conn{}        // output
	eb := &event.Bus{}         // both input and output

	// Domain
	upload.New(router, store, eb)
	retrieve.New(router, store)

	return http.ListenAndServe(":8080", router)
}
