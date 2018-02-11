package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/HackGT/SponsorshipPortal/backend/app"
	"github.com/HouzuoGuo/tiedot/db"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

type Login struct {
	*revel.Controller
}

var (
	mySigningKey = os.Getenv("signingKey")
	superToken   = os.Getenv("superToken")
)

// TODO: convert to SQL
func (c Login) Index() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	username := jsonData["username"].(string)
	password := jsonData["password"].(string)
	id, valid := IsValidUser(username, password)
	resp := make(map[string]interface{})
	if !valid {
		resp["error"] = "invalid user"
		return c.RenderJSON(resp)
	}
	createdToken, err := GetNewToken([]byte(mySigningKey), id)
	if err != nil {
		c.RenderError(err)
	}
	resp["token"] = createdToken
	return c.RenderJSON(resp)
}

// TODO: convert to SQL
func (c Login) AddUser() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	username := jsonData["username"].(string)
	password := jsonData["password"].(string)
	company := jsonData["company"].(string)
	token := jsonData["token"].(string)
	resp := make(map[string]interface{})
	if token != superToken {
		resp["error"] = "invalid token"
		return c.RenderJSON(resp)
	}
	sponsors := app.PortalDB.Use("Sponsors")
	var query interface{}
	json.Unmarshal([]byte(fmt.Sprintf(`[{"eq": "%s", "in": ["username"]}]`, username)), &query)
	queryResult := make(map[int]struct{})
	err := db.EvalQuery(query, sponsors, &queryResult)
	for range queryResult {
		resp["error"] = "user already exists"
		return c.RenderJSON(resp)
	}
	// insert user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	_, err = sponsors.Insert(map[string]interface{}{
		"company":  company,
		"username": username,
		"password": string(hashedPassword)})
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	resp["status"] = "done"
	return c.RenderJSON(resp)
}

// TODO: convert to SQL
func IsValidUser(username string, password string) (string, bool) {
	sponsors := app.PortalDB.Use("Sponsors")
	var query interface{}
	revel.INFO.Println(fmt.Sprintf(`[{"eq": "%s", "in": ["username"]}]`, username))
	json.Unmarshal([]byte(fmt.Sprintf(`[{"eq": "%s", "in": ["username"]}]`, username)), &query)
	queryResult := make(map[int]struct{})
	err := db.EvalQuery(query, sponsors, &queryResult)
	if err != nil {
		return "", false
	}
	for id := range queryResult {
		revel.INFO.Println(id)
		readBack, err := sponsors.Read(id)
		if err != nil {
			return "", false
		}
		err = bcrypt.CompareHashAndPassword([]byte(readBack["password"].(string)), []byte(password))
		if err == nil {
			return strconv.Itoa(id), true
		}
		return "", false
	}
	return "", false
}

func GetNewToken(mySigningKey []byte, id string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

func ParseToken(myToken string) (*jwt.Token, bool) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	if err == nil && token.Valid {
		return token, true
	}
	return nil, false
}
