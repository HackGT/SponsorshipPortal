package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host         string        `default:""`
	Port         int           `default:"3000"`
	WriteTimeout time.Duration `default:"15s"`
	ReadTimeout  time.Duration `default:"15s"`
	IdleTimeout  time.Duration `default:"60s"`
	ShutdownWait time.Duration `default:"15s"`
	Prod         bool          `default:"false"`
}

var config Config
var log = logrus.New()

func init() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse config from environment variables")
	}

	log.Infof("Initializing server with config: %+v", config)

	if config.Prod {
		log.SetLevel(logrus.InfoLevel)
		log.Formatter = &logrus.JSONFormatter{}
	} else {
		log.SetLevel(logrus.DebugLevel)
		log.Formatter = &logrus.TextFormatter{}
	}
}

func main() {
	// Parts of this is adapted from
	// https://github.com/gorilla/mux#graceful-shutdown

	addr := fmt.Sprintf("%s:%v", config.Host, config.Port)
	r := createRouter()
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      r,
	}

	// Start the server in a goroutine so it does not block
	go func() {
		log.Infof("Server started and listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.WithError(err).Fatal("Failed to start server")
		}
		log.Info("Server stopped")
	}()

	// Catch SIGINT's (Ctrl+C) and attempt a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a quit signal is received
	<-c

	log.Info("Shutdown signal received")
	ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownWait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Info("Shutting down")
	os.Exit(0)
}
