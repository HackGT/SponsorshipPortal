package auth

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

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
		w.Write([]byte(string(response) + "\nMalformed JSON Request"))
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonUser.Password), bcrypt.DefaultCost)
		if err != nil {
			w.Write([]byte("Unspecified error adding user."))
		} else {
			user.Create(u.db, jsonUser.Org_id, jsonUser.Email, string(hashedPassword))
			w.Write([]byte("Email: " + string(jsonUser.Email) + ", Password Hash: " + string(hashedPassword) + ", Org_Id: " + strconv.FormatInt(jsonUser.Org_id, 10) + ". User added to database."))
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Placeholder"))
}

func Load(r *mux.Router, db *sqlx.DB) {
	u := &userController{db}
	r.HandleFunc("", u.Create).Methods("PUT")
	r.HandleFunc("/login", Login).Methods("POST")
}
