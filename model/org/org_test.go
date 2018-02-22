package org

import (
	"database/sql"
	"testing"

	modeltesting "github.com/HackGT/SponsorshipPortal/model/testing"
)

func TestByID(t *testing.T) {
	getReqChan := make(chan *modeltesting.QueryRequest, 1)
	getRespChan := make(chan error, 1)
	db := &modeltesting.MockConnection{
		GetReq:  getReqChan,
		GetResp: getRespChan,
	}

	// success case
	var id int64 = 1
	getRespChan <- nil
	_, found, err := ByID(db, id)
	if err != nil || found != true {
		t.Errorf("Expected no error and found=true, instead got: found=%v, err=%v", found, err)
	}
	query := <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != id {
		t.Errorf("Expected query args to contain only id=%v, instead got: args=%v", id, query.Args)
	}

	// error - no rows
	id = int64(2)
	getRespChan <- sql.ErrNoRows
	_, found, err = ByID(db, id)
	if err != sql.ErrNoRows || found != false {
		t.Errorf("Expected error=sql.ErrNoRows and found=false, instead got: found=%v, err=%v", found, err)
	}
	query = <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != id {
		t.Errorf("Expected query args to contain only id=%v, instead got: args=%v", id, query.Args)
	}
}

func TestByName(t *testing.T) {
	getReqChan := make(chan *modeltesting.QueryRequest, 1)
	getRespChan := make(chan error, 1)
	db := &modeltesting.MockConnection{
		GetReq:  getReqChan,
		GetResp: getRespChan,
	}

	// success case
	name := "name1"
	getRespChan <- nil
	_, found, err := ByName(db, name)
	if err != nil || found != true {
		t.Errorf("Expected no error and found=true, instead got: found=%v, err=%v", found, err)
	}
	query := <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != name {
		t.Errorf("Expected query args to contain only name=%v, instead got: args=%v", name, query.Args)
	}

	// error - no rows
	name = "name2"
	getRespChan <- sql.ErrNoRows
	_, found, err = ByName(db, name)
	if err != sql.ErrNoRows || found != false {
		t.Errorf("Expected error=sql.ErrNoRows and found=false, instead got: found=%v, err=%v", found, err)
	}
	query = <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != name {
		t.Errorf("Expected query args to contain only name=%v, instead got: args=%v", name, query.Args)
	}
}

func TestCreate(t *testing.T) {
	execReqChan := make(chan *modeltesting.QueryRequest, 1)
	execRespChan := make(chan *modeltesting.SQLResultErrorPair, 1)
	db := &modeltesting.MockConnection{
		ExecReq:  execReqChan,
		ExecResp: execRespChan,
	}

	execRespChan <- &modeltesting.SQLResultErrorPair{}
	name := "org_name"
	_, err := Create(db, name)
	if err != nil {
		t.Errorf("Expected nil-error, instead got: error=%v", err)
	}
	query := <-execReqChan
	if len(query.Args) != 1 || query.Args[0] != name {
		t.Errorf("Expected query args to contain only name=%v, instead got: args=%v", name, query.Args)
	}
}
