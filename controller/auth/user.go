package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"

	"github.com/HackGT/SponsorshipPortal/model/user"
)

type User struct {
	Email    string
	Password string
	Org_id   int64
}

func Create(w http.ResponseWriter, r *http.Request) {
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
			execReqChan := make(chan *QueryRequest, 1)
			execRespChan := make(chan *SQLResultErrorPair, 1)
			db := &Connection{
				ExecReq:  execReqChan,
				ExecResp: execRespChan,
			}
			execRespChan <- &SQLResultErrorPair{}
			user.Create(db, jsonUser.Org_id, jsonUser.Email, string(hashedPassword))
			w.Write([]byte("Email: " + string(jsonUser.Email) + ", Password Hash: " + string(hashedPassword) + ", Org_Id: " + string(jsonUser.Org_id) + ". User added to database."))
		}
	}

}

func Load(r *mux.Router) {
	r.Methods("PUT").HandlerFunc(Create)
}
