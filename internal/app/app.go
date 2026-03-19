package app

import (
	"github.com/2Delight/mlist-backend/internal/config"
	"github.com/2Delight/mlist-backend/internal/server"
)

type App struct {
	server *server.Server
}

func New(config config.Config) *App {
	return &App{
		server: server.New(config.Port),
	}
}

func (a *App) Start() error {
	return a.server.Start()
}
