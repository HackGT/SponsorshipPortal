package controllers

import (
	"github.com/HackGT/SponsorshipPortal/backend/app"
	"github.com/blevesearch/bleve"
	"github.com/revel/revel"
)

type Search struct {
	*revel.Controller
}

func (c Search) Index() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	phrase := jsonData["q"].(string)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	_, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequest(query)
	searchResults, err := app.BleveIndex.Search(search)
	//originalDocs := getOriginalDocsFromSearchResults(searchResults, BleveIndex)
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["results"] = searchResults
	return c.RenderJSON(resp)
}
