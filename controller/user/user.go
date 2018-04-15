package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HackGT/SponsorshipPortal/model/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type userController struct {
	db *sqlx.DB
}

func (u *userController) SaveState(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t struct {
		State string `json:"state"`
	}
	err := decoder.Decode(&t)
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid request"))
		return
	}
	sponsorID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid id"))
		return
	}
	user, exist, err := user.ByID(u.db, int64(sponsorID))
	if err != nil || !exist {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot get user"))
		return
	}
	user.State = t.State
	_, err = user.Save(u.db)
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot save state"))
		return
	}
	resp := struct {
		Status string `json:"status"`
	}{"ok"}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func (u *userController) FetchState(w http.ResponseWriter, r *http.Request) {
	sponsorID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid id"))
		return
	}
	user, exist, err := user.ByID(u.db, int64(sponsorID))
	if err != nil || !exist {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot get user"))
		return
	}
	resp := struct {
		State string `json:"state"`
	}{user.State}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.WithError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func Load(r *mux.Router, db *sqlx.DB) {
	u := &userController{db}
	r.HandleFunc("/{id}/state", u.SaveState).Methods("POST")
	r.HandleFunc("/{id}/state", u.FetchState).Methods("GET")
}
