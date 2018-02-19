package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/HackGT/SponsorshipPortal/server"
)

func main() {
	// Parts of this is adapted from
	// https://github.com/gorilla/mux#graceful-shutdown

	app, err := server.New()
	if err != nil {
		panic(err)
	}
	log := app.Logger

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
