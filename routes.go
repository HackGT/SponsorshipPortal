package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/controller"
)

func NewRouter(logger *logrus.Logger) http.Handler {
	r := mux.NewRouter()

	// Load controllers
	controller.Load(r)

	// Add handler for static files
	r.Methods("GET").Handler(http.FileServer(http.Dir("./client/static/")))

	// Attach logging and recovery middlewares
	var handler http.Handler
	handler = handlers.LoggingHandler(logger.Writer(), r)
	handler = handlers.RecoveryHandler(handlers.RecoveryLogger(logger))(handler)
	return handler
}
