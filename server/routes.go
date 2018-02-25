package server

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	ctrl "github.com/HackGT/SponsorshipPortal/controller"
	"github.com/HackGT/SponsorshipPortal/controller/auth"
	"github.com/HackGT/SponsorshipPortal/logger"
)

func loggingMiddleware(out io.Writer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(out, next)
	}
}

func authMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return auth.AuthHandler(next)
	}
}

func (app *App) NewRouter() http.Handler {
	r := mux.NewRouter()

	// Load controllers
	controller := ctrl.New(app.DB, app.Config)
	controller.Load(r)

	// Add handler for static files
	r.Methods("GET").Handler(http.FileServer(http.Dir("./client/static/")))

	// Attach logging and recovery middlewares
	log := logger.New(app.Config)
	r.Use(authMiddleware())
	r.Use(loggingMiddleware(log.Writer()))
	r.Use(handlers.RecoveryHandler(handlers.RecoveryLogger(log)))
	return r
}
