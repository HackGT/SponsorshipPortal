package auth

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"	
)

type testHandler struct {
}

func (t testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SUCCESS"))
}

func generateECKeyPair() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("openssl", "ecparam", "-name", "secp521r1", "-out", "./secp521r1.pem")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.WithError(errors.New(stderr.String())).Warn(out.String())
		log.WithError(err).Error("Unable to generate secp521r1 EC specifications file.")
	}
	cmd = exec.Command("openssl", "ecparam", "-in", "./secp521r1.pem", "-genkey", "-noout", "-out", "./ecprivatekey.pem")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.WithError(errors.New(stderr.String())).Warn(out.String())
		log.WithError(err).Error("Unable to generate private key.")
	}
	cmd = exec.Command("openssl", "ec", "-in", "./ecprivatekey.pem", "-pubout", "-out", "./ecpublickey.pem")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.WithError(errors.New(stderr.String())).Warn(out.String())
		log.WithError(err).Error("Unable to generate public key.")
	}
	ecPrivateKeyRaw, _ := ioutil.ReadFile("./ecprivatekey.pem")
	ecPublicKeyRaw, _ := ioutil.ReadFile("./ecpublickey.pem")
	os.Setenv("EC_PRIVATE_KEY", string(ecPrivateKeyRaw))
	os.Setenv("EC_PUBLIC_KEY", string(ecPublicKeyRaw))
}

func cleanPEMFiles() {
	os.Remove("./secp521r1.pem")
	os.Remove("./ecprivatekey.pem")
	os.Remove("./ecpublickey.pem")
	os.Unsetenv("EC_PRIVATE_KEY");
	os.Unsetenv("EC_PUBLIC_KEY");
}

func TestRequireAuthenticationSuccess(t *testing.T) {
	//initialize required resources
	logger, hook := test.NewNullLogger()
	ra := reqAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//success case
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(expires))
	claims.SetIssuedAt(timeNow)
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow)
	claims.SetSubject("test-success")
	claims.Set("eid", "testrequireauthsuccess@hack.gt")
	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Error("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Error("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-success")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	ra.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		log.WithError(hook.LastEntry().Data["error"].(error)).Error("Internal Error found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "SUCCESS"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestRequireAuthenticationNoJWT(t *testing.T) {
	//initialize required resources
	logger, hook := test.NewNullLogger()
	ra := reqAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"

	//fail case - no jwt
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-fail-no-jwt")))
	w := httptest.NewRecorder()
	ra.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "401 Unauthorized\n"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	//test error log
	expectedError := jws.ErrNoTokenInRequest
	if hook.LastEntry().Data["error"].(error) != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : hook.LastEntry().Data["error"].(error)}).Error("Unexpected error found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestRequireAuthenticationExpired(t *testing.T) {
	//initialize required resources
	logger, hook := test.NewNullLogger()
	ra := reqAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//fail case - expired jwt
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(-2 * expires))
	claims.SetIssuedAt(timeNow.Add(-3 * expires))
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow.Add(-3 * expires))
	claims.SetSubject("test-fail-expired")
	claims.Set("eid", "testrequireauthexpired@hack.gt")
	jwtObj := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Warn("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwtObj.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-fail-expired")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	ra.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "401 Unauthorized\n"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	//test error log
	expectedError := jwt.ErrTokenIsExpired
	if hook.LastEntry().Data["error"].(error) != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : hook.LastEntry().Data["error"].(error)}).Error("Unexpected error found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestRequireAuthenticationInvalid(t *testing.T) {
	//initialize required resources
	logger, hook := test.NewNullLogger()
	ra := reqAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//fail case - invalid jwt
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(expires))
	claims.SetIssuedAt(timeNow)
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow)
	claims.SetSubject("test-fail-invalid")
	claims.Set("eid", "testrequireauthinvalid@hack.gt")
	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Warn("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-fail-invalid")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	//generate new keys for validation, invalidate old keys
	generateECKeyPair()
	ra.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "401 Unauthorized\n"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	//test error log
	expectedError := crypto.ErrECDSAVerification
	if hook.LastEntry().Data["error"].(error) != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : hook.LastEntry().Data["error"].(error)}).Error("Unexpected error found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestRequireNoAuthenticationSuccessNoJWT(t *testing.T) {
	//initialize required resources
	logger, _ := test.NewNullLogger()
	rna := reqNoAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"

	//success case - no jwt
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-success-no-jwt")))
	w := httptest.NewRecorder()
	//generate new keys for validation, invalidate old keys
	generateECKeyPair()
	rna.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "SUCCESS"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestNoRequireAuthenticationSuccessExpired(t *testing.T) {
	//initialize required resources
	logger, _ := test.NewNullLogger()
	rna := reqNoAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//success case - expired
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(-2 * expires))
	claims.SetIssuedAt(timeNow.Add(-3 * expires))
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow.Add(-3 * expires))
	claims.SetSubject("test-success-expired")
	claims.Set("eid", "testrequirenoauthexpired@hack.gt")
	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Warn("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-success-expired")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	rna.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "SUCCESS"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestNoRequireAuthenticationSuccessInvalid(t *testing.T) {
	//initialize required resources
	logger, _ := test.NewNullLogger()
	rna := reqNoAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//success case - invalid jwt
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(expires))
	claims.SetIssuedAt(timeNow)
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow)
	claims.SetSubject("test-success-invalid")
	claims.Set("eid", "testrequirenoauthinvalid@hack.gt")
	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Warn("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-success-invalid")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	//generate new keys for validation, invalidate old keys
	generateECKeyPair()
	rna.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "SUCCESS"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	cleanPEMFiles()
}

func TestNoRequireAuthenticationFail(t *testing.T) {
	//initialize required resources
	logger, _ := test.NewNullLogger()
	rna := reqNoAuthHandler{handler : testHandler{}, log : logger}
	generateECKeyPair()

	host := "localhost:3000"
	expires := 15 * time.Minute

	//fail case - valid jwt
	timeNow := time.Now()
	claims := jws.Claims{}
	claims.SetAudience(host)
	claims.SetExpiration(timeNow.Add(expires))
	claims.SetIssuedAt(timeNow)
	claims.SetIssuer(host)
	claims.SetNotBefore(timeNow)
	claims.SetSubject("test-fail-valid-jwt")
	claims.Set("eid", "testrequirenoauthfail@hack.gt")
	jwt := jws.NewJWT(claims, crypto.SigningMethodES512)
	rawPrivateKey := os.Getenv("EC_PRIVATE_KEY")
	if rawPrivateKey == "" {
		log.Error("Error: EC Private Key env var not set. Did you generate your EC key pair?")
		t.Fail()
	}
	privateKey, err := crypto.ParseECPrivateKeyFromPEM([]byte(rawPrivateKey))
	if err != nil {
		log.WithError(err).Warn("Unable to parse private key.")
		t.Fail()
	}
	token, err := jwt.Serialize(privateKey)
	if err != nil {
		log.WithError(err).Warn("Unable to serialize JWT.")
		t.Fail()
	}
	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test-fail-valid-jwt")))
	testRequest.Header.Set("Authorization", "Bearer " + string(token))
	w := httptest.NewRecorder()
	rna.ServeHTTP(w, testRequest)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned response
	expectedResponse := "401 Unauthorized\n"
	if string(w.Body.Bytes()) != expectedResponse {
		log.WithFields(log.Fields{"Expected" : expectedResponse, "Retrieved" : string(w.Body.Bytes())}).Error("Unexpected HTTP response found.")
		t.Fail()
	}
	cleanPEMFiles()
}
