package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"

	"github.com/HackGT/SponsorshipPortal/config"
)

type reqAuthHandler struct {
	handler     http.Handler
	log         *logrus.Logger
	authConfig  *config.AuthenticationConfig
}
type reqNoAuthHandler struct {
	handler     http.Handler
	log         *logrus.Logger
	authConfig  *config.AuthenticationConfig
}

func RequireAuthentication(logger *logrus.Logger, authConfig *config.AuthenticationConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return reqAuthHandler{next, logger, authConfig}
	}
}

func RequireNoAuthentication(logger *logrus.Logger, authConfig *config.AuthenticationConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return reqNoAuthHandler{next, logger, authConfig}
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
	validator := jws.NewValidator(nil, a.authConfig.EXPLeeway, a.authConfig.NBFLeeway, nil)
	rawPublicKey := os.Getenv("EC_PUBLIC_KEY")
	if rawPublicKey == "" {
		a.log.Error("Error: EC Public Key env var not set. Did you generate your EC key pair?")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM([]byte(rawPublicKey))
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
	validator := jws.NewValidator(nil, na.authConfig.EXPLeeway, na.authConfig.NBFLeeway, nil)
	rawPublicKey := os.Getenv("EC_PUBLIC_KEY")
	if rawPublicKey == "" {
		na.log.Error("Error: EC Public Key env var not set. Did you generate your EC key pair?")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM([]byte(rawPublicKey))
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
