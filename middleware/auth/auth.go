package auth

import (
	"time"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
)

//Leeways for token expiration and token not-before times. Default to 5 minutes
var expLeeway int64 = 300000000000
var nbfLeeway int64 = 300000000000

type authHandler struct {
	handler http.Handler
}

func RequireAuthentication() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return AuthHandler(next)
	}
}

func (a authHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	token, err := jws.ParseJWTFromRequest(req)
	if err != nil {
		//JWT Parse error - no token found, probably not authenticated
		http.Redirect(w, req, "localhost:3000/user/login", 302)
		return
	}
	validator := jws.NewValidator(nil, time.Duration(expLeeway), time.Duration(nbfLeeway), nil)
	err = token.Validate([]byte("test"), crypto.SigningMethodES512, validator)
	if err != nil {
		http.Redirect(w, req, "localhost:3000/user/login", 302) //Would issue 401 unauthorized but redirect only allows 3xx codes
		return
	}
	a.handler.ServeHTTP(w, req)
}

func AuthHandler(h http.Handler) http.Handler {
	return authHandler{h}
}
