package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/HackGT/SponsorshipPortal/controller"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Load controllers
	controller.Load(r)

	// Add handler for static files
	r.Methods("GET").Handler(http.FileServer(http.Dir("./static/")))

	// Attach logging and recovery middlewares
	var handler http.Handler
	handler = handlers.LoggingHandler(log.Writer(), r)
	handler = handlers.RecoveryHandler(handlers.RecoveryLogger(log))(handler)
	return handler
}
