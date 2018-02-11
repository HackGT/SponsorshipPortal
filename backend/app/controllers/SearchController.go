package controllers

import (
	"fmt"
	"strconv"

	"github.com/HackGT/SponsorshipPortal/backend/app"
	"github.com/blevesearch/bleve"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
)

type Search struct {
	*revel.Controller
}

// TODO: convert to SQL
func (c Search) Index() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	phrase := jsonData["query"].(string)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	token, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequest(query)
	searchResults, err := app.BleveIndex.Search(search)
	resp := make(map[string]interface{})
	if err == nil {
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
		resultUUID := ""
		for _, element := range searchResults.Hits {
			resultUUID += element.ID
		}
		fmt.Println(fmt.Sprintf(`{ 
			hackgtmetricsversion: 1,
			serviceName: "sponsorshipportal-hackgt4-search",
			values: {value:1},
			tags: {search: "%s", user: "%s", resultNum: %d, resultUuid: "%s"}
		`, phrase, readBack["company"], searchResults.Total, resultUUID))
	}
	resp["error"] = err
	resp["results"] = searchResults
	return c.RenderJSON(resp)
}
