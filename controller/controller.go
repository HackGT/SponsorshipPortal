package controller

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/config"
	"github.com/HackGT/SponsorshipPortal/controller/auth"
	"github.com/HackGT/SponsorshipPortal/controller/health"
	"github.com/HackGT/SponsorshipPortal/controller/sample"
)

type Controller struct {
	DB     *sqlx.DB
	Config *config.Config
	Log    *logrus.Logger
}

func New(db *sqlx.DB, config *config.Config, log *logrus.Logger) *Controller {
	return &Controller{
		DB:     db,
		Config: config,
		Log:	log,
	}
}

func (c *Controller) Load(r *mux.Router) {
	// Register controllers and their respective path prefixes
	health.Load(r.PathPrefix("/_health").Subrouter(), c.DB.DB, c.Config.Database)
	sample.Load(r.PathPrefix("/sample").Subrouter())
	auth.Load(r.PathPrefix("/user").Subrouter(), c.DB, c.Log, c.Config.Auth)
}
