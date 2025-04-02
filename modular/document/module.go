package document

import (
	"modular/common/infra/db"
	"modular/common/infra/event"
	"modular/common/module"
	"modular/document/api"

	"github.com/julienschmidt/httprouter"
)

var _ module.Interface = (*Module)(nil)

type Module struct {
	name        string
	db          *db.Conn
	eventBus    *event.Bus
	httpHandler *api.HTTPHandler
}

// NewModule creates a new document module
func NewModule(name string, db *db.Conn, eventBus *event.Bus) *Module {
	return &Module{
		name:     name,
		db:       db,
		eventBus: eventBus,
	}
}

// Name returns the name of the module
func (m *Module) Name() string {
	return m.name
}

// Init initializes the module
func (m *Module) Init() error {
	// TODO: Initialize the module

	return nil
}

// Start starts the module
func (m *Module) Start() error {
	// Start services and infrastructure
	return nil
}

// Stop stops the module
func (m *Module) Stop() error {
	// Stop the services and infrastructure
	return nil
}

// GetHTTPHandler returns the HTTP handler for the module
func (m *Module) GetHTTPHandler() *api.HTTPHandler {
	return m.httpHandler
}

func (m *Module) SetupRoutes(router *httprouter.Router) {
	m.httpHandler.RegisterRoutes(router)
}

// RegisterEventHandlers here to satisfy module.Interface
func (*Module) RegisterEventHandlers() {
	// Register event handlers here
}
