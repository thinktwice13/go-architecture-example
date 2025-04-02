package bootstrap

import (
	"net/http"
	"vertical/feature/retrieve"
	"vertical/feature/upload"
	"vertical/shared/db"
	"vertical/shared/event"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	router := httprouter.New() // input
	store := &db.Conn{}        // output
	eb := &event.Bus{}         // both input and output

	// Domain
	upload.New(router, store, eb)
	retrieve.New(router, store)

	return http.ListenAndServe(":8080", router)
}
