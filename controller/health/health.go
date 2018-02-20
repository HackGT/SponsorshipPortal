package health

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/database"
)

type healthController struct {
	db     *sql.DB
	config *config.DatabaseConfig
}

func (h *healthController) Health(w http.ResponseWriter, r *http.Request) {
	ready, err := database.IsReadyWithInstance(h.db, h.config)
	var status int
	var msg string
	if err != nil {
		status = http.StatusInternalServerError
		msg = "Internal error"
		log.WithError(err).Warn("Error retrieving health (migration state)")
	} else if ready {
		status = http.StatusOK
		msg = "OK"
	} else {
		status = http.StatusServiceUnavailable
		msg = "Unhealthy"
	}
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func Load(r *mux.Router, db *sql.DB, config *config.DatabaseConfig) {
	h := &healthController{db, config}
	r.Methods("HEAD", "GET").HandlerFunc(h.Health)
}
