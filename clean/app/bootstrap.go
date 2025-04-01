package app

import (
	"clean/adapter/logging"
	"clean/adapter/router"
	"clean/infra/config"
	"clean/infra/server"
	"clean/infra/storage"
)

// Application represents the application
type Application struct {
	config *config.Config
	logger *logging.DefaultLogger
	server *server.Server
	db     *storage.DB
}

func New() *Application {
	cfg := config.Load()
	log := &logging.DefaultLogger{}

	return &Application{
		config: cfg,
		logger: log,
	}
}

func (a *Application) Init() error {
	a.db = &storage.DB{}
	container := NewContainer(a.logger, a.db)
	router := router.NewRouter(container.DocumentHandler)
	httpRouter := router.Routes()
	a.server = server.NewServer(httpRouter, a.logger)
	return nil
}

func (a *Application) start() error {
	_ = a.Init()
	return a.server.Start()
}

func Run() error {
	app := New()
	return app.start()
}
