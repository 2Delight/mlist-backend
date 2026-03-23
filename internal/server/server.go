package server

import (
	"fmt"
	"net/http"
	"time"
)

const (
	timeout = 1 * time.Second
)

type ModelsProvider interface{}

type Server struct {
	server         http.Server
	modelsProvider ModelsProvider
}

func New(port uint16, modelsProvider ModelsProvider) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf(":%d", port)
	s.server.Handler = s.setRouter()
	s.modelsProvider = modelsProvider
	return s
}

func (s *Server) setRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /ping", wrapHandlerFunc(
		s.pingHandler,
		addTimeout,
		// addLogging,
	))
	return mux
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
