package main

import (
	"os"

	"github.com/2Delight/mlist-backend/internal/app"
	"github.com/2Delight/mlist-backend/internal/config"
)

func main() {
	configPath := os.Args[1]
	conf, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	app, err := app.New(conf)
	if err != nil {
		panic(err)
	}

	panic(app.Start())
}
