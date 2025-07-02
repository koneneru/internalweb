package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"sync"
	"time"

	"atsinfoService/internal/data"
	"atsinfoService/internal/jsonlog"

	_ "github.com/alexbrainman/odbc"
)

type database struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

type config struct {
	port int
	env  string
	dbd  database
	dbs  database
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 9000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.dbs.dsn, "dbs-dsn", "", "MSSQL SVET data source name")
	flag.IntVar(&cfg.dbs.maxOpenConns, "dbs-max-open-conns", 25, "MSSQL SVET max open connections")
	flag.IntVar(&cfg.dbs.maxIdleConns, "dbs-max-idle-conns", 25, "MSSQL SVET max idle connections")
	flag.DurationVar(&cfg.dbs.maxIdleTime, "dbs-max-idle-time", 15*time.Minute, "MSSQL SVET max commection idle time")

	flag.StringVar(&cfg.dbd.dsn, "dbd-dsn", "", "MSSQL DIRECTUM data source name")
	flag.IntVar(&cfg.dbd.maxOpenConns, "dbd-max-open-conns", 25, "MSSQL DIRECTUM max open connections")
	flag.IntVar(&cfg.dbd.maxIdleConns, "dbd-max-idle-conns", 25, "MSSQL DIRECTUM max idle connections")
	flag.DurationVar(&cfg.dbd.maxIdleTime, "dbd-max-idle-time", 15*time.Minute, "MSSQL DIRECTUM max commection idle time")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	var databases [2]*sql.DB
	for i, d := range []database{cfg.dbs, cfg.dbd} {
		db, err := openDB(d)
		if err != nil {
			logger.PrintFatal(err, nil)
			os.Exit(0)
		}
		defer db.Close()
		databases[i] = db
		logger.PrintInfo("database connection pool established", nil)
	}

	app := application{
		config: cfg,
		logger: logger,
		models: data.NewModels(databases[0], databases[1]),
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(dbi database) (*sql.DB, error) {
	db, err := sql.Open("odbc", dbi.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbi.maxOpenConns)
	db.SetMaxIdleConns(dbi.maxIdleConns)
	db.SetConnMaxIdleTime(dbi.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
