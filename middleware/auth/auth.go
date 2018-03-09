package auth

import (
	"time"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
)

//Leeways for token expiration and token not-before times. Default to 5 minutes
const expLeeway time.Duration = 3 * time.Minute
const nbfLeeway time.Duration = 3 * time.Minute

type reqAuthHandler struct {
	handler http.Handler
}
type reqNoAuthHandler struct {
	handler http.Handler
}

func RequireAuthentication() mux.MiddlewareFunc {
	return ReqAuthHandler
}

func RequireNoAuthentication() mux.MiddlewareFunc {
	return ReqNoAuthHandler
}

func (a reqAuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token, err := jws.ParseJWTFromRequest(req)
	if err != nil {
		//JWT Parse error - no token found, probably not authenticated
		log.WithError(err).Warn("Unauthorized access attempt - no token.")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	validator := jws.NewValidator(nil, expLeeway, nbfLeeway, nil)
	rawPublicKey, err := ioutil.ReadFile("./ecpublickey.pem")
	if err != nil {
		log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = token.Validate(publicKey, crypto.SigningMethodES512, validator)
	if err != nil {
		//Invalid JWT
		log.WithError(err).Warn("Unauthorized access attempt - invalid token.")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	//Authorized
	a.handler.ServeHTTP(w, req)
}

func (na reqNoAuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token, err := jws.ParseJWTFromRequest(req)
        if err != nil {
                //JWT Parse error - no token found, probably not authenticated
		na.handler.ServeHTTP(w, req)
		return
        }
        validator := jws.NewValidator(nil, expLeeway, nbfLeeway, nil)
	rawPublicKey, err := ioutil.ReadFile("./ecpublickey.pem")
        if err != nil {
                log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
        }
        publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
        if err != nil {
                log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
        }
        err = token.Validate(publicKey, crypto.SigningMethodES512, validator)
        if err != nil {
                //Invalid JWT
                na.handler.ServeHTTP(w, req)
                return
        }
        //Authorized
        log.Warn("Attempted to create an account or login while authenticated.")
	http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
}

func ReqAuthHandler(next http.Handler) http.Handler {
	return reqAuthHandler{next}
}
func ReqNoAuthHandler(next http.Handler) http.Handler {
	return reqNoAuthHandler{next}
}
