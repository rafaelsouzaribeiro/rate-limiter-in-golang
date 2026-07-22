package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Mux      *http.ServeMux
	addr     string
	handlers map[string]http.HandlerFunc
}

func NewServer(addr string) *Server {
	return &Server{
		Mux:      http.NewServeMux(),
		addr:     addr,
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (s *Server) RegisterHandler(path string, handler http.HandlerFunc) {
	s.handlers[path] = handler
}

func (s *Server) Start() error {
	for path, handler := range s.handlers {
		s.Mux.Handle(path, handler)
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", s.addr), s.Mux)
}
