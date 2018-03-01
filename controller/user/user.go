package user

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type userController struct {
	db *sqlx.DB
}

func (u *userController) SaveState(w http.ResponseWriter, r *http.Request) {
	sponsorId := context.Get(r, "id")

}

func (u *userController) FetchState(w http.ResponseWriter, r *http.Request) {

}

func Load(r *mux.Router, db *sqlx.DB) {
	u := &userController{db}
	r.HandleFunc("/{id}/state", u.SaveState).Methods("POST")
	r.HandleFunc("/{id}/state", u.SaveState).Methods("GET")
}
