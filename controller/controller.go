package controller

import (
	"github.com/gorilla/mux"

	"github.com/HackGT/SponsorshipPortal/controller/sample"
)

func Load(r *mux.Router) {

	// Register controllers and their respective path prefixes
	sample.Load(r.PathPrefix("/sample").Subrouter())
}
