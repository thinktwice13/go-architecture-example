package document

import (
	"modular/common/infra/db"
	"modular/common/infra/event"
	"modular/common/module"
	"modular/document/api"
	"modular/document/application"
	"modular/document/repo"

	"github.com/julienschmidt/httprouter"
)

var _ module.Interface = (*Module)(nil)

type Module struct {
	name     string
	db       *db.Conn
	eventBus *event.Bus
	router   *httprouter.Router
}

// NewModule creates a new document module
func NewModule(name string, db *db.Conn, eventBus *event.Bus, router *httprouter.Router) *Module {
	return &Module{name, db, eventBus, router}
}

func (m *Module) Name() string {
	return m.name
}

func (m *Module) Start() error {
	docRepo := repo.NewRepo(m.db)
	svc := application.NewDocumentService(docRepo, m.eventBus)
	apiHandler := api.NewHTTPHandler(svc)
	apiHandler.RegisterRoutes(m.router)
	return nil
}
