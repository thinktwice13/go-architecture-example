package server

import (
	"net/http"
)

type Server struct {
	router http.Handler
}

func NewServer(router http.Handler) *Server {
	return &Server{router}
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
