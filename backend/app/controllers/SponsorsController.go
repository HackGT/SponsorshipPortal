package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/HackGT/SponsorshipPortal/backend/app"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

type Sponsors struct {
	*revel.Controller
}

func (c Sponsors) Save() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	state := jsonData["state"].(string)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	token, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	sponsors := app.PortalDB.Use("Sponsors")
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return c.RenderError(err)
	}
	readBack, err := sponsors.Read(id)
	if err != nil {
		return c.RenderError(err)
	}
	readBack["state"] = state
	err = sponsors.Update(id, readBack)
	if err != nil {
		return c.RenderError(err)
	}
	res["status"] = "saved"
	return c.RenderJSON(res)
}

func (c Sponsors) Load() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	token, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	sponsors := app.PortalDB.Use("Sponsors")
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return c.RenderError(err)
	}
	readBack, err := sponsors.Read(id)
	if err != nil {
		return c.RenderError(err)
	}
	if val, ok := readBack["state"]; ok {
		res["state"] = val
		return c.RenderJSON(res)
	}
	res["status"] = "none"
	return c.RenderJSON(res)
}

func (c Sponsors) FlagParticipant() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	tokenStr := jsonData["token"].(string)
	participantID := jsonData["participant_id"].(string)
	res := make(map[string]interface{})
	token, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	sponsors := app.PortalDB.Use("Sponsors")
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return c.RenderError(err)
	}
	readBack, err := sponsors.Read(id)
	if err != nil {
		return c.RenderError(err)
	}
	if val, ok := readBack["met"]; ok {
		var met []string
		json.Unmarshal([]byte(val.(string)), &met)
		met = append(met, participantID)
		metJSON, _ := json.Marshal(met)
		readBack["met"] = string(metJSON)
	} else {
		var s []string
		s = append(s, participantID)
		metJSON, err := json.Marshal(s)
		if err != nil {
			return c.RenderError(err)
		}
		readBack["met"] = string(metJSON)
	}
	err = sponsors.Update(id, readBack)
	if err != nil {
		return c.RenderError(err)
	}
	res["status"] = "done"
	return c.RenderJSON(res)
}

func (c Sponsors) PeopleMet() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	token, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	sponsors := app.PortalDB.Use("Sponsors")
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return c.RenderError(err)
	}
	readBack, err := sponsors.Read(id)
	if err != nil {
		return c.RenderError(err)
	}
	var met []string
	json.Unmarshal([]byte(readBack["met"].(string)), &met)
	res["people"] = met
	return c.RenderJSON(res)
}
