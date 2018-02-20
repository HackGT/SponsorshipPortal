package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/mattes/migrate"
	log "github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/database"
	"github.com/HackGT/SponsorshipPortal/logger"
	"github.com/HackGT/SponsorshipPortal/server"
)

var runMigrations = flag.Bool("migrate", false, "run database migrations and exit")

func dbMigrate() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(conf).WithFields(log.Fields{
		"host":   conf.Database.Host,
		"user":   conf.Database.User,
		"dbname": conf.Database.DbName,
	})

	log.Info("Migrating database...")

	err = database.Migrate(conf.Database.ConnectionString)
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Warn("No changes detected.")
		} else {
			log.WithError(err).Fatal("Failed to migrate database")
		}
	}

	log.Info("Finished migrating.")
	os.Exit(0)
}

func startServer() {
	// Parts of this is adapted from
	// https://github.com/gorilla/mux#graceful-shutdown

	app, err := server.New()
	if err != nil {
		panic(err)
	}

	// Start the server in a goroutine so it does not block
	go func() {
		log.Infof("Server started and listening on %s", app.Config.Server.Addr())
		if err := app.Server.ListenAndServe(); err != nil {
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
	ctx, cancel := context.WithTimeout(context.Background(), app.Config.Server.ShutdownWait)
	defer cancel()

	app.Server.Shutdown(ctx)

	log.Info("Shutting down")
	os.Exit(0)
}

func main() {
	flag.Parse()
	if *runMigrations {
		dbMigrate()
	} else {
		startServer()
	}
}
