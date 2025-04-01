package server

import (
	"clean/adapter/logging"
	"fmt"
	"net/http"
)

// Server represents an HTTP server
type Server struct {
	router http.Handler
	logger logging.Logger
}

// NewServer creates a new server
func NewServer(router http.Handler, logger logging.Logger) *Server {
	return &Server{
		router: router,
		logger: logger,
	}
}

// Start starts the server
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", 8080)
	s.logger.Log("Starting server on " + addr)
	return http.ListenAndServe(addr, s.router)
}
