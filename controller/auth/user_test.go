package auth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	httptest "net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"golang.org/x/crypto/bcrypt"	
)

type testUser struct {
	Email string
	Password string
	Org_id int
}

type testLoginUser struct {
	Email string
	Password string
}

func createSQLiteDatabaseConnection(file string) (*sqlx.DB) {
	os.MkdirAll("./testing-user", 0755)
	os.Create("./testing-user/test-" + file + ".db")
	db, err := sqlx.Open("sqlite3", "./testing-user/test-" + file + ".db")
	if err != nil {
		log.WithError(err).Error("Unable to open sqlite database.")
		return nil
	}
	createTables(db)
	return db
}

func createTables(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS sponsor_orgs (
			id integer PRIMARY KEY AUTOINCREMENT,
			name text NOT NULL,
			created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			deleted_at timestamptz
	);`)
	if err != nil {
		log.WithError(err).Fatal("Unable to create table sponsor_orgs")
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS sponsors (
			id integer PRIMARY KEY AUTOINCREMENT,
			org_id integer REFERENCES sponsor_orgs(id),
			email text UNIQUE NOT NULL,
			password text NOT NULL,
			created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			deleted_at timestamptz
	);`)
	if err != nil {
		log.WithError(err).Fatal("Unable to create table sponsors")
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS participants (
			id integer PRIMARY KEY AUTOINCREMENT,
			registration_id text UNIQUE NOT NULL,
			document tsvector,
			created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
			deleted_at timestamptz
	);`)
	if err != nil {
		log.WithError(err).Fatal("Unable to create table participants")
	}
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
}

func cleanPEMFiles() {
	os.Remove("./secp521r1.pem")
	os.Remove("./ecprivatekey.pem")
	os.Remove("./ecpublickey.pem")
}

func cleanSQLiteDatabase(file string, db *sqlx.DB) {
	db.Close()
	os.RemoveAll("./testing-user/")
}

func validateResponseJWT(t *testing.T, w *httptest.ResponseRecorder, host string) {
	type TokenResp struct {
		Token string
	}
	var parsedToken TokenResp
	err := json.Unmarshal(w.Body.Bytes(), &parsedToken)
	if err != nil {
		log.WithError(err).Error("Unable to unmarshall JWT to struct.")
		t.Fail()
	}
	pseudoJWTRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test")))
	pseudoJWTRequest.Header.Set("Authorization", "Bearer " + parsedToken.Token)
	token, err := jws.ParseJWTFromRequest(pseudoJWTRequest)
	if err != nil {
		//JWT Parse error - no token found, probably not authenticated
		log.WithError(err).Warn("Unauthorized access attempt - no token.")
		t.Fail()
	}
	validator := jws.NewValidator(nil, 1, 1, nil)
	rawPublicKey, err := ioutil.ReadFile("./ecpublickey.pem")
	if err != nil {
		log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
		t.Fail()
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
		t.Fail()
	}
	err = token.Validate(publicKey, crypto.SigningMethodES512, validator)
	if err != nil {
		//Invalid JWT
		log.WithError(err).Warn("Unauthorized access attempt - invalid token.")
		t.Fail()
	}
}

func validateJWT(t *testing.T, jwt []byte, host string) {
	type TokenResp struct {
		Token string
	}
	parsedToken := TokenResp{string(jwt)}
	pseudoJWTRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test")))
	pseudoJWTRequest.Header.Set("Authorization", "Bearer " + parsedToken.Token)
	token, err := jws.ParseJWTFromRequest(pseudoJWTRequest)
	if err != nil {
		//JWT Parse error - no token found, probably not authenticated
		log.WithError(err).Warn("Unauthorized access attempt - no token.")
		t.Fail()
	}
	validator := jws.NewValidator(nil, 1, 1, nil)
	rawPublicKey, err := ioutil.ReadFile("./ecpublickey.pem")
	if err != nil {
		log.WithError(err).Error("Error reading EC public key. Are you sure you generated your EC public-private key pair?")
		t.Fail()
	}
	publicKey, err := crypto.ParseECPublicKeyFromPEM(rawPublicKey)
	if err != nil {
		log.WithError(err).Error("Error parsing ECDSA public key from file. Are you sure you have the correct format? It should be ES512.")
		t.Fail()
	}
	err = token.Validate(publicKey, crypto.SigningMethodES512, validator)
	if err != nil {
		//Invalid JWT
		log.WithError(err).Warn("Unauthorized access attempt - invalid token.")
		t.Fail()
	}
}

func TestCreateUserSuccess(t *testing.T) {
	testName := "create-user-success"

	//initialize required resources
	logger, _ := test.NewNullLogger()
	db := createSQLiteDatabaseConnection(testName)
	u := userController{db : db, log : logger}
	generateECKeyPair()

	//create reference sponsor_org
	_, err := db.Exec("INSERT INTO sponsor_orgs (id, name) VALUES (?, ?)", 5, "Microsoft")
	if err != nil {
		log.WithError(err).Error("Error inserting reference sponsor_org into the database.")
		t.Fail()	
	}

	host := "localhost:3000"
	

	//success case
	expectedUser := testUser{Email : "testnewemail@hack.gt", Password : "test", Org_id : 5}
	req, _ := json.Marshal(expectedUser)
	r := httptest.NewRequest("PUT", host + "/user", bytes.NewReader(req))
	w := httptest.NewRecorder()
	u.Create(w, r)
	//test response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response code.")
		t.Fail()
	}
	//test returned JWT
	validateResponseJWT(t, w, host)
	//test database
	rows, err := db.Query("SELECT Email, Password, Org_id FROM sponsors WHERE Email=?", expectedUser.Email)
	if err != nil {
		log.WithError(err).Error("Unable to retrieve rows from database.")
		t.Fail()
	} else {
		var retrievedUser testUser
		rows.Next()
		err = rows.Scan(&retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Org_id)
		if err != nil {
			log.WithError(err).Error("Unable to retrieve columns from returned row.")
			t.Fail()
		}
		if retrievedUser.Email != expectedUser.Email {
			log.WithFields(log.Fields{"Expected" : expectedUser.Email, "Retrieved" : retrievedUser.Email}).Error("Email was not as expected.")
			t.Fail()
		}
		if retrievedUser.Password == expectedUser.Password {
			log.Error("Password was not hashed.")
			t.Fail()
		}
		if retrievedUser.Org_id != expectedUser.Org_id {
			log.WithFields(log.Fields{"Expected" : expectedUser.Org_id, "Retrieved" : retrievedUser.Org_id}).Error("Org_id was not as expected.")
			t.Fail()
		}
		rows.Close()
	}

	//clean up
	cleanSQLiteDatabase(testName, db)
	cleanPEMFiles()
}

func TestCreateUserFail(t *testing.T) {
	testName := "create-user-fail"

	//initialize required resources
	logger, hook := test.NewNullLogger()
	db := createSQLiteDatabaseConnection(testName)
	u := userController{db : db, log : logger}
	generateECKeyPair()

	//create reference sponsor_org
	_, err := db.Exec("INSERT INTO sponsor_orgs (id, name) VALUES (?, ?)", 5, "Microsoft")
	if err != nil {
		log.WithError(err).Error("Error inserting reference sponsor_org into the database.")
		t.Fail()	
	}

	host := "localhost:3000"
	
	//fail case
	existingUser := testUser{Email : "testexistingemail@hack.gt", Password : "notapplicable", Org_id : 5}
	_, err = db.Exec("INSERT INTO sponsors (Email, Password, Org_id) VALUES (?, ?, ?)", existingUser.Email, existingUser.Password, existingUser.Org_id)
	if err != nil {
		log.WithError(err).Error("Error inserting reference user into the database.")
		t.Fail()	
	}
	req, _ := json.Marshal(existingUser)
        r := httptest.NewRequest("PUT", host + "/user", bytes.NewReader(req))
        w := httptest.NewRecorder()
	expectedError := "UNIQUE constraint failed: sponsors.email"
	u.Create(w, r)
	retrievedError := hook.LastEntry().Data["error"].(error).Error()
	if retrievedError != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : retrievedError}).Error("Received unexpected error.")
		t.Fail()
	}

	//clean up
	cleanSQLiteDatabase(testName, db)
	cleanPEMFiles()
}

func TestLoginSuccess(t *testing.T) {
	testName := "login-success"

	//initialize required resources
	logger, hook := test.NewNullLogger()
	db := createSQLiteDatabaseConnection(testName)
	u := userController{db : db, log : logger}
	generateECKeyPair()

	//create reference sponsor_org
	_, err := db.Exec("INSERT INTO sponsor_orgs (id, name) VALUES (?, ?)", 5, "Microsoft")
	if err != nil {
		log.WithError(err).Error("Error inserting reference sponsor_org into the database.")
		t.Fail()	
	}

	host := "localhost:3000"
	
	//success case
	existingUser := testUser{Email : "testloginsuccess@hack.gt", Password : "testlogin", Org_id : 5}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(existingUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.WithError(err).Error("Error generating password hash.")
		t.Fail()
	}
	existingUser.Password = string(hashedPassword)
	_, err = db.Exec("INSERT INTO sponsors (org_id, email, password) VALUES (?, ?, ?)", existingUser.Org_id, existingUser.Email, existingUser.Password)
	if err != nil {
		log.WithError(err).Error("Error inserting reference user into the database.")
		t.Fail()	
	}

	loggingInUser := testLoginUser{Email : "testloginsuccess@hack.gt", Password: "testlogin"}
	req, _ := json.Marshal(loggingInUser)
	r := httptest.NewRequest("PUT", host + "/user", bytes.NewReader(req))
        w := httptest.NewRecorder()
	u.Login(w, r)
	//test response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response code.")
		log.WithError(hook.LastEntry().Data["error"].(error)).Error("Last log message: " + hook.LastEntry().Message)
		t.Fail()
	}
	//test returned JWT
	validateResponseJWT(t, w, host)

	//clean up
	cleanSQLiteDatabase(testName, db)
	cleanPEMFiles()
}

func TestLoginNoSuchUser(t *testing.T) {
	testName := "login-no-such-user"

	//initialize required resources
	logger, hook := test.NewNullLogger()
	db := createSQLiteDatabaseConnection(testName)
	u := userController{db : db, log : logger}
	generateECKeyPair()

	//create reference sponsor_org
	_, err := db.Exec("INSERT INTO sponsor_orgs (id, name) VALUES (?, ?)", 5, "Microsoft")
	if err != nil {
		log.WithError(err).Error("Error inserting reference sponsor_org into the database.")
		t.Fail()	
	}

	host := "localhost:3000"

	//fail case - nonexistent user
	nonexistentUser := testLoginUser{Email : "testnosuchuser@hack.gt", Password: "testlogin"}
	req, _ := json.Marshal(nonexistentUser)
	r := httptest.NewRequest("PUT", host + "/user/login", bytes.NewReader(req))
        w := httptest.NewRecorder()
	u.Login(w, r)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test log error
	expectedError := sql.ErrNoRows
	retrievedError := hook.LastEntry().Data["error"].(error)
	if retrievedError != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : retrievedError}).Error("Received unexpected error.")
		t.Fail()
	}

	//clean up
	cleanSQLiteDatabase(testName, db)
	cleanPEMFiles()
}

func TestLoginIncorrect(t *testing.T) {
	testName := "login-incorrect"

	//initialize required resources
	logger, hook := test.NewNullLogger()
	db := createSQLiteDatabaseConnection(testName)
	u := userController{db : db, log : logger}
	generateECKeyPair()

	//create reference sponsor_org
	_, err := db.Exec("INSERT INTO sponsor_orgs (id, name) VALUES (?, ?)", 5, "Microsoft")
	if err != nil {
		log.WithError(err).Error("Error inserting reference sponsor_org into the database.")
		t.Fail()	
	}

	host := "localhost:3000"

	existingUser := testUser{Email : "testloginincorrect@hack.gt", Password : "testlogin", Org_id : 5}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(existingUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.WithError(err).Error("Error generating password hash.")
		t.Fail()
	}
	existingUser.Password = string(hashedPassword)
	_, err = db.Exec("INSERT INTO sponsors (org_id, email, password) VALUES (?, ?, ?)", existingUser.Org_id, existingUser.Email, existingUser.Password)
	if err != nil {
		log.WithError(err).Error("Error inserting reference user into the database.")
		t.Fail()	
	}

	//fail case - incorrect password
	wrongPasswordUser := testLoginUser{Email : "testloginincorrect@hack.gt", Password: "wrongpassword"}
	req, _ := json.Marshal(wrongPasswordUser)
	r := httptest.NewRequest("PUT", host + "/user/login", bytes.NewReader(req))
        w := httptest.NewRecorder()
	u.Login(w, r)
	//test http response code
	if w.Code != http.StatusUnauthorized {
		log.WithFields(log.Fields{"Expected" : http.StatusUnauthorized, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test log error
	expectedError := bcrypt.ErrMismatchedHashAndPassword
	retrievedError := hook.LastEntry().Data["error"].(error)
	if retrievedError != expectedError {
		log.WithFields(log.Fields{"Expected" : expectedError, "Retrieved" : retrievedError}).Error("Received unexpected error.")
		t.Fail()
	}

	//clean up
	cleanSQLiteDatabase(testName, db)
	cleanPEMFiles()
}

func TestCreateJWT(t *testing.T) {
	generateECKeyPair()

	host := "localhost:3000"
	
	//success case
	token, err := CreateJWT("testcreatejwt@hack.gt", host, host)
	if err != nil {
		log.WithError(err).Error("Failed to create JWT.")
		t.Fail()
		return
	}
	//test JWT validation
	validateJWT(t, token, host)

	cleanPEMFiles()
}

func TestReTokenSuccess(t *testing.T) {
	//initialize required resources
	logger, _ := test.NewNullLogger()
	u := userController{db : nil, log : logger}
	generateECKeyPair()

	host := "localhost:3000"

	testRequest := httptest.NewRequest("GET", host, bytes.NewReader([]byte("test")))
	w := httptest.NewRecorder()
	u.ReToken(w, testRequest)
	//test http response code
	if w.Code != http.StatusOK {
		log.WithFields(log.Fields{"Expected" : http.StatusOK, "Retrieved" : w.Code}).Error("Unexpected HTTP response status found.")
		t.Fail()
	}
	//test returned JWT
	validateResponseJWT(t, w, host)	

	cleanPEMFiles()
}
