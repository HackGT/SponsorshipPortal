package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	ctrl "github.com/HackGT/SponsorshipPortal/controller"
	"github.com/HackGT/SponsorshipPortal/logger"
)

func (app *App) NewRouter() http.Handler {
	r := mux.NewRouter()

	// Load controllers
	controller := ctrl.New(app.DB, app.Config)
	controller.Load(r)

	// Add handler for static files
	r.Methods("GET").Handler(http.FileServer(http.Dir("./client/static/")))

	// Attach logging and recovery middlewares
	log := logger.New(app.Config)
	var handler http.Handler
	handler = handlers.LoggingHandler(log.Writer(), r)
	handler = handlers.RecoveryHandler(handlers.RecoveryLogger(log))(handler)
	return handler
}
