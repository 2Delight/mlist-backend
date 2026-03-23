package main

import (
	"context"
	"os"

	"github.com/2Delight/mlist-backend/internal/app"
	"github.com/2Delight/mlist-backend/internal/config"
	"github.com/2Delight/mlist-backend/internal/logger"
)

func main() {
	configPath := os.Args[1]
	conf, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	logger.Setup(conf.Logger)
	logger.Info(context.Background(), "logger is locked and loaded")

	app, err := app.New(conf)
	if err != nil {
		panic(err)
	}

	panic(app.Start())
}
