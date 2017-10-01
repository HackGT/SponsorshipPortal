package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

var BleveIndex bleve.Index

type ResumeParser struct {
	*revel.Controller
}

type SearchEngine struct {
	*revel.Controller
}

type Participants struct {
	*revel.Controller
}

type Sponsors struct {
	*revel.Controller
}

type ParseResume struct {
	resumeURL     string
	participantID string
}

func (p ParseResume) Run() {
	revel.INFO.Println(p.resumeURL)
	out, err := exec.Command("ruby", "/home/brow/Documents/textextract.rb", p.resumeURL).Output()
	if err != nil {
		revel.ERROR.Println(err)
		return
	}
	//revel.INFO.Println(string(out))
	data := struct {
		Resume string
	}{
		Resume: string(out),
	}

	// index some data
	BleveIndex.Index(p.participantID, data)
}

func (c ResumeParser) Index() revel.Result {
	resumeURL := c.Params.Query.Get("resume")
	participantID := c.Params.Query.Get("participant_id")
	jobs.Now(ParseResume{resumeURL, participantID})
	resp := make(map[string]interface{})
	resp["error"] = nil
	resp["status"] = "done"
	return c.RenderJSON(resp)
}

func (c SearchEngine) Index() revel.Result {
	phrase := c.Params.Query.Get("q")
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequest(query)
	searchResults, err := BleveIndex.Search(search)
	//originalDocs := getOriginalDocsFromSearchResults(searchResults, BleveIndex)
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["results"] = searchResults
	return c.RenderJSON(resp)
}

func (c Participants) Index() revel.Result {
	body := strings.NewReader(`q=["all"]`)
	req, err := http.NewRequest("POST", "http://localhost:8080/query?col=Participants", body)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.RenderError(err)
	}
	in := bodyBytes
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	return c.RenderJSON(raw)
}

func (c Sponsors) Save() revel.Result {
	sponsorID := c.Params.Query.Get("sponsor_id")
	state := c.Params.Query.Get("state")
	body := strings.NewReader(`doc={"state": ` + state + `}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/update?col=Sponsors&id="+sponsorID, body)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	defer resp.Body.Close()

	res := make(map[string]interface{})
	res["status"] = 1
	return c.RenderJSON(res)
}

func (c Sponsors) Load() revel.Result {
	sponsorID := c.Params.Query.Get("sponsor_id")
	resp, err := http.Get("http://localhost:8080/get?col=Sponsors&id=" + sponsorID)
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.RenderError(err)
	}
	in := bodyBytes
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	return c.RenderJSON(raw)
}
