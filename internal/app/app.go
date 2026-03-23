package app

import (
	"github.com/2Delight/mlist-backend/internal/config"
	"github.com/2Delight/mlist-backend/internal/metrics"
	"github.com/2Delight/mlist-backend/internal/providers/models"
	"github.com/2Delight/mlist-backend/internal/server"
	"github.com/2Delight/mlist-backend/internal/storage"
)

type Providers struct {
	Models models.Provider
}

type App struct {
	server    *server.Server
	providers Providers
	storage   storage.Storage
}

func New(config config.Config) (*App, error) {
	storage, err := storage.New(config.DB.GetDBURL(), config.DB.MigrationsPath)
	if err != nil {
		return nil, err
	}

	modelsProvider := models.NewProvider(storage)

	metrics.SetupMetrics(config.Metrics.Port)
	return &App{
		server: server.New(config.Server.Port, modelsProvider),
		providers: Providers{
			Models: modelsProvider,
		},
		storage: storage,
	}, nil
}

func (a *App) Start() error {
	return a.server.Start()
}
