package auth

import (
	"time"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"

	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
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

type AuthUser struct {
	Email string
	Password string
}

type authResponse struct {
	Token string
}

func (u userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonUser User
	response, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(response, &jsonUser)
	if err != nil {
		log.WithError(err).Warn("Error while unmarshalling json request.")
		http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonUser.Password), bcrypt.DefaultCost)
	if err != nil {
 		log.WithError(err).Warn("Error while generating password hash.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = user.Create(u.db, jsonUser.Org_id, jsonUser.Email, string(hashedPassword))
	if err != nil {
		log.WithError(err).Warn("Error adding user to database.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Debug("Email: " + string(jsonUser.Email) + ", Password Hash: " + string(jsonUser.Password) + ", OrgID: " + strconv.FormatInt(jsonUser.Org_id, 10) + ". User added to database.")
	//send JWT back
	serializedJWT, err := CreateJWT(jsonUser.Email, r.Host)
        if err != nil { 
                log.WithError(err).Warn("Failed to create JWT.") 
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return 
        }
        ar := authResponse{Token: string(serializedJWT)}
        token, err := json.Marshal(ar)
        if err != nil {
                log.WithError(err).Warn("Error marshalling json web token.")
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
		log.WithError(err).Warn("Error while unmarshalling json request")
                http.Error(w, "Malformed JSON Request", http.StatusBadRequest)
		return
	}
	validUser, exists, err := user.ByEmail(u.db, jsonUser.Email)
	if err != nil || !exists {
		if !exists {
			log.WithError(err).Warn("Somebody attempted to log into an account that doesn't exist: " + jsonUser.Email)
		} else {
			log.WithError(err).Warn("Error retrieving email: " + jsonUser.Email + " from the database.")
		}
		http.Error(w, "Login Error", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(validUser.Password), []byte(jsonUser.Password))
	if err != nil {
		log.WithError(err).Warn("Somebody attempted to log into account: " + jsonUser.Email)
		http.Error(w, "Login Error", http.StatusUnauthorized)
		return
	}
	//Logged in, send back JWT
	serializedJWT, err := CreateJWT(jsonUser.Email, r.Host)
	if err != nil {
		log.WithError(err).Warn("Failed to create JWT.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	ar := authResponse{Token: string(serializedJWT)}
	token, err := json.Marshal(ar)
	if err != nil {
		log.WithError(err).Warn("Error marshalling json web token.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(token)
}

func ReToken(w http.ResponseWriter, req *http.Request) {
	email := req.Header.Get("eid")
        serializedJWT, err := CreateJWT(email, req.Host)
        if err != nil { 
                log.WithError(err).Warn("Failed to create JWT.") 
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return 
        }
        ar := authResponse{Token: string(serializedJWT)}
        returnToken, err := json.Marshal(ar)
        if err != nil {
                log.WithError(err).Warn("Error marshalling json web token.")
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


func CreateJWT(email string, host string) ([]byte, error) {
	t := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(t.Add(expires))
	claims.SetIssuedAt(t)
	claims.SetIssuer(host + "/user/login")
	jwtid := sha256.Sum256([]byte(email + strconv.Itoa(count)))
	encodedJwtid := make([]byte, base64.StdEncoding.EncodedLen(len(jwtid[:])))
	base64.StdEncoding.Encode(encodedJwtid, jwtid[:])
	claims.SetJWTID(string(encodedJwtid))
	claims.SetNotBefore(t)
	claims.SetSubject(subject)
	claims.Set("eid", email)

	log.Debug("JWT Claims:")
	log.Debug("Audience: " + host)
        log.Debug("Expiration Time: " + t.Add(expires).String())
        log.Debug("Issued At: " + t.String())
        log.Debug("Issuer: " + host + "/user/login")
        log.Debug("Base64 JWTID: " + string(encodedJwtid))
        log.Debug("NotBefore: " + t.String())
        log.Debug("Subject: " + subject)

	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey, err := ioutil.ReadFile("./ecprivatekey.pem")
	if err != nil {
		log.WithError(err).Error("Error reading EC private key. Are you sure you have generated your EC public-private key pair?")
		return nil, err
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM(rawPrivateKey)
	if err != nil {
		log.WithError(err).Error("Error parsing ECDSA private key from file. Are you sure you have the correct format? It should be ES512.")
		return nil, err
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize jwt.")
		return nil, err
	}
	count++
	return token, nil
}

func Load(r *mux.Router, db *sqlx.DB) {
	u := &userController{db}
	createUserSubR := r.PathPrefix("").Subrouter()
	createUserSubR.Methods("PUT").HandlerFunc(u.Create)
	createUserSubR.Use(auth.RequireNoAuthentication())
	loginSubR := r.PathPrefix("/login").Subrouter()
	loginSubR.Methods("POST").HandlerFunc(u.Login)
	loginSubR.Use(auth.RequireNoAuthentication())
	renewSubR := r.PathPrefix("/renew").Subrouter()
	renewSubR.Methods("GET").HandlerFunc(ReToken)
	renewSubR.Use(auth.RequireAuthentication())
}
