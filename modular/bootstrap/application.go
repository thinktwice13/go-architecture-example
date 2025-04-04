package boostrap

import (
	"hexagonal/core/config"
	"modular/common/infra/db"
	"modular/common/infra/event"
	"modular/common/module"
	"modular/document"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RunApplication() error {
	// Infra
	_ = config.Load()
	dbConn := &db.Conn{}
	eb := &event.Bus{}
	router := httprouter.New()

	// Modules
	modules := []module.Interface{
		document.NewModule("document", dbConn, eb, router),
		// processing.NewModule(....)
	}

	// Start modules
	for _, mod := range modules {
		_ = mod.Start()
	}

	// Start server
	return http.ListenAndServe(":8080", router)
}
