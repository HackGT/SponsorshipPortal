package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
)

var Server *http.Server
var Config *config.Config
var Logger *logrus.Logger

var log *logrus.Logger

func init() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	log = logrus.New()
	log.Infof("Initializing server with config: %+v", config)
	if config.Prod {
		log.SetLevel(logrus.InfoLevel)
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.SetLevel(logrus.DebugLevel)
		log.Formatter = &logrus.TextFormatter{}
	}

	r := NewRouter()
	server := &http.Server{
		Addr: config.Addr(),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      r,
	}

	Server = server
	Config = config
	Logger = log

}
