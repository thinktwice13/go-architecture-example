package app

import (
	"clean/adapter/api/handler"
	"clean/adapter/eventbus"
	"clean/adapter/logging"
	"clean/adapter/persistence"
	"clean/infra/storage"
	"clean/usecase/document"
)

// Container holds all application dependencies
// This centralizes dependency injection and makes it easier to manage
type Container struct {
	Logger          *logging.DefaultLogger
	DB              *storage.DB
	EventBus        *eventbus.EventBus
	DocumentRepo    *persistence.DocumentRepo
	UploadUseCase   *document.UploadUseCase
	RetrieveUseCase *document.RetrieveUseCase
	DocumentHandler *handler.DocumentHandler
}

func NewContainer(logger *logging.DefaultLogger, db *storage.DB) *Container {
	c := &Container{
		Logger: logger,
		DB:     db,
	}

	c.EventBus = eventbus.NewEventBus(logger)
	c.DocumentRepo = persistence.NewDocumentRepo(db, logger)
	c.UploadUseCase = document.NewUploadUseCase(c.DocumentRepo, c.EventBus, logger)
	c.RetrieveUseCase = document.NewRetrieveUseCase(c.DocumentRepo, c.EventBus, logger)
	c.DocumentHandler = handler.NewDocumentHandler(c.UploadUseCase, c.RetrieveUseCase)
	return c
}
