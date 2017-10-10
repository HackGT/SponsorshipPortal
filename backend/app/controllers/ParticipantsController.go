package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/HackGT/SponsorshipPortal/backend/app"
	"github.com/HackGT/SponsorshipPortal/backend/app/portaljobs"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

type Participants struct {
	*revel.Controller
}

func (c Participants) Index() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	_, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	participants := app.PortalDB.Use("Participants")
	var query interface{}
	json.Unmarshal([]byte(fmt.Sprintf(`["all"]`)), &query)
	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, participants, &queryResult); err != nil {
		return c.RenderError(err)
	}
	for id := range queryResult {
		readBack, err := participants.Read(id)
		if err != nil {
			return c.RenderError(err)
		}
		if err == nil {
			res[strconv.Itoa(id)] = readBack
		}
	}
	return c.RenderJSON(res)
}

func (c Participants) Add() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	name := jsonData["name"].(string)
	email := jsonData["email"].(string)
	resumeID := jsonData["resumeId"].(string)
	token := jsonData["token"].(string)
	resp := make(map[string]interface{})
	if token != superToken {
		resp["error"] = "invalid token"
		return c.RenderJSON(resp)
	}
	participants := app.PortalDB.Use("Participants")
	// insert participant
	id, err := participants.Insert(map[string]interface{}{
		"name":     name,
		"email":    email,
		"resumeId": resumeID})
	if err != nil {
		revel.ERROR.Println(err)
		return c.RenderError(err)
	}
	startJob(resumeID, strconv.Itoa(id), name, email)
	resp["status"] = "done"
	return c.RenderJSON(resp)
}

func startJob(resumeID string, id string, name string, email string) string {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("registration-dev-uploads"),
		Key:    aws.String(resumeID),
	})
	urlStr, err := req.Presign(time.Minute)
	if err != nil {
		revel.ERROR.Println(err)
	}
	jobs.Now(portaljobs.ParseResume{ResumeURL: urlStr, ParticipantID: id, Name: name, Email: email})
	return urlStr
}

func (c Participants) ViewResume() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	resumeID := jsonData["resumeId"].(string)
	tokenStr := jsonData["token"].(string)
	res := make(map[string]interface{})
	_, valid := ParseToken(tokenStr)
	if !valid {
		res["error"] = "invalid token"
		return c.RenderJSON(res)
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("registration-dev-uploads"),
		Key:    aws.String(resumeID),
	})
	urlStr, err := req.Presign(time.Minute)
	if err != nil {
		return c.RenderError(err)
	}
	res["fileURL"] = urlStr
	return c.RenderJSON(res)
}
