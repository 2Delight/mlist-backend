package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/2Delight/mlist-backend/internal/providers/models"
)

const (
	timeout = 1 * time.Second
)

type ModelsProvider interface {
	GetModels(ctx context.Context) ([]models.Model, error)
	CreateModel(ctx context.Context, model models.Model) (models.Model, error)
	UpdateModel(ctx context.Context, modelID int, model models.Model) (models.Model, error)
	DeleteModel(ctx context.Context, modelID int) error
	LookupModel(ctx context.Context, repository string, version string) (bool, error)
}

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
		addMetrics,
		addLogging,
	))
	mux.Handle("GET /models", wrapHandlerFunc(
		s.getModelsHandler,
		addTimeout,
		addMetrics,
		addLogging,
	))
	mux.Handle("PUT /models", wrapHandlerFunc(
		s.updateModelHandler,
		addTimeout,
		addMetrics,
		addLogging,
	))
	mux.Handle("DELETE /models", wrapHandlerFunc(
		s.deleteModelHandler,
		addTimeout,
		addMetrics,
		addLogging,
	))
	mux.Handle("GET /lookup-model", wrapHandlerFunc(
		s.lookupModelHandler,
		addTimeout,
		addMetrics,
		addLogging,
	))
	return mux
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
