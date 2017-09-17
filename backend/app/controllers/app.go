package controllers

import (
	"os/exec"

	"github.com/blevesearch/bleve"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

type ResumeParser struct {
	*revel.Controller
}

type ParseResume struct {
	resumeURL string
}

func (p ParseResume) Run() {
	out, err := exec.Command("ruby", "/home/brow/Documents/textextract.rb").Output()
	if err != nil {
		revel.ERROR.Println(err)
	}
	bleveIndex, err := bleve.Open("example.bleve")
	if err != nil {
		revel.ERROR.Println(err)
	}
	data := struct {
		Resume string
	}{
		Resume: string(out),
	}

	// index some data
	bleveIndex.Index(p.resumeURL, data)
}

func (c ResumeParser) Index() revel.Result {
	resumeURL := c.Params.Query.Get("resume")
	jobs.Now(ParseResume{resumeURL})
	resp := make(map[string]interface{})
	resp["error"] = nil
	resp["status"] = "done"
	return c.RenderJSON(resp)
}
