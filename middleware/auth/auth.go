package auth

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

//Leeways for token expiration and token not-before times. Default to 3 minutes
const expLeeway time.Duration = 3 * time.Minute
const nbfLeeway time.Duration = 3 * time.Minute

type reqAuthHandler struct {
	handler http.Handler
	log	*logrus.Logger
}
type reqNoAuthHandler struct {
	handler http.Handler
	log	*logrus.Logger
}

func RequireAuthentication(logger *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return reqAuthHandler{next, logger}
	}
}

func RequireNoAuthentication(logger *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return reqNoAuthHandler{next, logger}
	}
}

func (a reqAuthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token, err := jws.ParseJWTFromRequest(req)
	if err != nil {
		//JWT Parse error - no token found, probably not authenticated
		a.log.WithError(err).Warn("Unauthorized access attempt - no token.")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	validator := jws.NewValidator(nil, expLeeway, nbfLeeway, nil)
	rawPublicKey, err := ioutil.ReadFile("./ecpublickey.pem")
	if err != nil {
		a.log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		a.log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = token.Validate(publicKey, crypto.SigningMethodES512, validator)
	if err != nil {
		//Invalid JWT
		a.log.WithError(err).Warn("Unauthorized access attempt - invalid token.")
		http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
		return
	}
	//Authorized
	eid := token.Claims().Get("eid").(string)
	req.Header.Add("eid", eid)
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
		na.log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		na.log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
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
	na.log.Warn("Attempted to create an account or login while authenticated.")
	http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
}
