package server

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/database"
	"github.com/HackGT/SponsorshipPortal/logger"
)

type App struct {
	Server *http.Server
	Config *config.Config
	DB     *sqlx.DB
}

func New() (*App, error) {
	app := &App{}

	config, err := config.Load()
	if err != nil {
		return nil, err
	}
	app.Config = config
	logger.SetGlobalLogger(config)

	log.Infof("Connecting to database with config: %+v", config.Database)
	db, err := database.New(config.Database)
	if err != nil {
		log.WithError(err).Warn("Failed to initialize database connection")
		return nil, err
	}
	app.DB = db

	log.Infof("Initializing server with config: %+v", config.Server)
	r := app.NewRouter()
	server := &http.Server{
		Addr: config.Server.Addr(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
		Handler:      r,
	}
	app.Server = server

	return app, nil
}
