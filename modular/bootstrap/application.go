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
		document.NewModule("document", dbConn, eb),
		// processing.NewModule(db, eventBus),
	}

	for _, mod := range modules {
		_ = mod.Start()
	}

	// Register routes
	for _, mod := range modules {
		mod.SetupRoutes(router)
	}

	// Register event handlers
	for _, mod := range modules {
		mod.RegisterEventHandlers()
	}

	// Start server
	return http.ListenAndServe(":8080", router)
}
