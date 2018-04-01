package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/HackGT/SponsorshipPortal/middleware/auth"
	"github.com/HackGT/SponsorshipPortal/model/user"
)

type userController struct {
	db *sqlx.DB
	log *logrus.Logger
}

type User struct {
	Email    string
	Password string
	Org_id   int64
}

type AuthUser struct {
	Email    string
	Password string
}

type JWToken string

type authResponse struct {
	Token JWToken
}

func (u userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonUser User
	response, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(response, &jsonUser)
	if err != nil {
		u.log.WithError(err).Warn("Error while unmarshalling json request.")
		http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonUser.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.WithError(err).Warn("Error while generating password hash.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = user.Create(u.db, jsonUser.Org_id, jsonUser.Email, string(hashedPassword))
	if err != nil {
		u.log.WithError(err).Warn("Error adding user to database.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	u.log.Debug("Email: " + string(jsonUser.Email) + ", Password Hash: " + string(jsonUser.Password) + ", OrgID: " + strconv.FormatInt(jsonUser.Org_id, 10) + ". User added to database.")
	//send JWT back
	serializedJWT, err := CreateJWT(jsonUser.Email, r.Host, "/user")
	if err != nil {
		u.log.WithError(err).Warn("Failed to create JWT.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	ar := authResponse{Token: JWToken(serializedJWT)}
	token, err := json.Marshal(ar)
	if err != nil {
		u.log.WithError(err).Warn("Error marshalling json web token.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(token)
}

func (u userController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonUser AuthUser
	response, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(response, &jsonUser)
	if err != nil {
		u.log.WithError(err).Warn("Error while unmarshalling json request")
		http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
		return
	}
	validUser, exists, err := user.ByEmail(u.db, jsonUser.Email)
	if err != nil || !exists {
		if !exists {
			u.log.WithError(err).Warn("Somebody attempted to log into an account that doesn't exist: " + jsonUser.Email)
		} else {
			u.log.WithError(err).Warn("Error retrieving email: " + jsonUser.Email + " from the database.")
		}
		http.Error(w, "Login Error", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(jsonUser.Password))
	if err != nil {
		u.log.WithError(err).Warn("Somebody attempted to log into account: " + jsonUser.Email)
		http.Error(w, "Login Error", http.StatusUnauthorized)
		return
	}
	//Logged in, send back JWT
	serializedJWT, err := CreateJWT(jsonUser.Email, r.Host, "/user/login")
	if err != nil {
		u.log.WithError(err).Warn("Failed to create JWT.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	ar := authResponse{Token: JWToken(serializedJWT)}
	token, err := json.Marshal(ar)
	if err != nil {
		u.log.WithError(err).Warn("Error marshalling json web token.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(token)
}

func (u userController) ReToken(w http.ResponseWriter, req *http.Request) {
	email := req.Header.Get("eid")
	serializedJWT, err := CreateJWT(email, req.Host, "/user/renew")
	if err != nil {
		u.log.WithError(err).Warn("Failed to create JWT.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	ar := authResponse{Token: JWToken(serializedJWT)}
	returnToken, err := json.Marshal(ar)
	if err != nil {
		u.log.WithError(err).Warn("Error marshalling json web token.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(returnToken)
}

//Here are the claims that will be used for the JWT
const expires time.Duration = 15 * time.Minute
const subject string = "auth"

//JWTID is sha256 hash of the user's email concatenated with a salt
var count int = 0 //Integer counter that ensures the jwtid is unique

func CreateJWT(email string, host string, issuer string) ([]byte, error) {
	t := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(t.Add(expires))
	claims.SetIssuedAt(t)
	claims.SetIssuer(host + issuer)
	jwtid := sha256.Sum256([]byte(email + strconv.Itoa(count)))
	encodedJwtid := make([]byte, base64.StdEncoding.EncodedLen(len(jwtid[:])))
	base64.StdEncoding.Encode(encodedJwtid, jwtid[:])
	claims.SetJWTID(string(encodedJwtid))
	claims.SetNotBefore(t)
	claims.SetSubject(subject)
	claims.Set("eid", email)

	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey, err := ioutil.ReadFile("./ecprivatekey.pem")
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM(rawPrivateKey)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		return nil, err
	}
	count++
	return token, nil
}

func Load(r *mux.Router, db *sqlx.DB, log *logrus.Logger) {
	u := &userController{db, log}
	createUserSubR := r.PathPrefix("").Methods("PUT").Subrouter()
	createUserSubR.PathPrefix("").HandlerFunc(u.Create)
	createUserSubR.Use(auth.RequireNoAuthentication(log))
	loginSubR := r.PathPrefix("/login").Methods("POST").Subrouter()
	loginSubR.PathPrefix("").HandlerFunc(u.Login)
	loginSubR.Use(auth.RequireNoAuthentication(log))
	renewSubR := r.PathPrefix("/renew").Methods("GET").Subrouter()
	renewSubR.PathPrefix("").HandlerFunc(u.ReToken)
	renewSubR.Use(auth.RequireAuthentication(log))
}
