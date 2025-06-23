package main

import (
	"flag"
	"os"
	"sync"

	"internalwebFrontend/internal/jsonlog"
)

type config struct {
	port int
	env  string
	api  string
}

type application struct {
	config config
	logger *jsonlog.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.api, "api", "localhost:4000", "Host of API")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
