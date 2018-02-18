package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/logger"
)

type App struct {
	Server *http.Server
	Config *config.Config
	Logger *logrus.Logger
}

func New() *App {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(config)

	log.Infof("Initializing server with config: %+v", config)

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
		Logger: log,
	}
}
