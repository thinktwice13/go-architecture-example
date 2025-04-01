// Package router provides the HTTP router for the application.
package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"clean/adapter/api/handler"
)

type Router struct {
	documentHandler *handler.DocumentHandler
}

func NewRouter(documentHandler *handler.DocumentHandler) *Router {
	return &Router{
		documentHandler: documentHandler,
	}
}

func (r *Router) Routes() *httprouter.Router {
	router := httprouter.New()

	// Document routes
	router.POST("/documents", r.documentHandler.Upload)
	router.GET("/documents/:id", r.documentHandler.Get)

	// Health check
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}
