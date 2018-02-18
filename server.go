package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/database"
	"github.com/HackGT/SponsorshipPortal/logger"
)

type App struct {
	Server *http.Server
	Config *config.Config
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func New() (*App, error) {
	config, err := config.Load()
	if err != nil {
		return nil, err
	}

	log := logger.New(config)

	log.Infof("Connecting to database with config: %+v", config.Database)
	db, err := database.New(config.Database)
	if err != nil {
		log.WithError(err).Warn("Failed to initialize database connection")
		return nil, err
	}

	log.Infof("Initializing server with config: %+v", config.Server)
	r := NewRouter(log)
	server := &http.Server{
		Addr: config.Server.Addr(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
		Handler:      r,
	}

	return &App{
		Server: server,
		Config: config,
		DB:     db,
		Logger: log,
	}, nil
}
