package auth

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/middleware/auth"
	"github.com/HackGT/SponsorshipPortal/model/user"
)

type userController struct {
	db *sqlx.DB
}

type User struct {
	Email    string
	Password string
	Org_id   int64
}

func (u userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonUser User
	response, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(response, &jsonUser)
	if err != nil {
		log.WithError(err).Warn("Error while unmarshalling json request.")
		http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonUser.Password), bcrypt.DefaultCost)
		if err != nil {
			log.WithError(err).Warn("Error while generating password hash.")
			http.Error(w, "Unspecified error while adding user.", http.StatusInternalServerError)
		} else {
			user.Create(u.db, jsonUser.Org_id, jsonUser.Email, string(hashedPassword))
			log.Debug("Email: " + string(jsonUser.Email) + ", Password Hash: " + string(hashedPassword) + ", Org_Id: " + strconv.FormatInt(jsonUser.Org_id, 10) + ". User added to database.")
			//Currently unimplemented, send JWT back
		}
	}

}

func (u userController) Login(w http.ResponseWriter, r *http.Request) {
	//Placeholder
}

func Load(r *mux.Router, db *sqlx.DB) {
	u := &userController{db}
	r.HandleFunc("", u.Create).Methods("PUT")
	r.HandleFunc("/login", u.Login).Methods("POST")
	r.Use(auth.RequireNoAuthentication())
}
