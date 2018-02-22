package user

import (
	"database/sql"
	"testing"

	modeltesting "github.com/HackGT/SponsorshipPortal/model/testing"
)

func TestByEmail(t *testing.T) {
	getReqChan := make(chan *modeltesting.QueryRequest, 1)
	getRespChan := make(chan error, 1)
	db := &modeltesting.MockConnection{
		GetReq:  getReqChan,
		GetResp: getRespChan,
	}

	// success case
	email := "testworkingemail@hack.gt"
	getRespChan <- nil
	_, found, err := ByEmail(db, email)
	if err != nil || found != true {
		t.Errorf("Expected no error and found=true, instead got: found=%v, err=%v", found, err)
	}
	query := <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != email {
		t.Errorf("Expected query args to contain only email=%v, instead got: args=%v", email, query.Args)
	}

	// error - no rows
	email = "testfailingemail@hack.gt"
	getRespChan <- sql.ErrNoRows
	_, found, err = ByEmail(db, email)
	if err != sql.ErrNoRows || found != false {
		t.Errorf("Expected error=sql.ErrNoRows and found=false, instead got: found=%v, err=%v", found, err)
	}
	query = <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != email {
		t.Errorf("Expected query args to contain only email=%v, instead got: args=%v", email, query.Args)
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
	orgID := int64(0)
	email := "test@hack.gt"
	password := "bloop"
	_, err := Create(db, orgID, email, password)
	if err != nil {
		t.Errorf("Expected nil-error, instead got: error=%v", err)
	}
	query := <-execReqChan
	if len(query.Args) != 3 || query.Args[0] != orgID || query.Args[1] != email || query.Args[2] != password {
		t.Errorf("Expected query args to contain [%v, %v, %v], instead got: args=%v", orgID, email, password, query.Args)
	}
}

func TestSave(t *testing.T) {
	execReqChan := make(chan *modeltesting.QueryRequest, 1)
	execRespChan := make(chan *modeltesting.SQLResultErrorPair, 1)
	db := &modeltesting.MockConnection{
		ExecReq:  execReqChan,
		ExecResp: execRespChan,
	}

	execRespChan <- &modeltesting.SQLResultErrorPair{}

	testUser := &User{
		ID:       int64(0),
		OrgID:    int64(1),
		Email:    "test@hack.gt",
		Password: "bloop",
	}

	_, err := testUser.Save(db)
	if err != nil {
		t.Errorf("Expected nil-error, instead got: error=%v", err)
	}
	query := <-execReqChan
	if len(query.Args) != 4 || query.Args[0] != testUser.OrgID || query.Args[1] != testUser.Email || query.Args[2] != testUser.Password || query.Args[3] != testUser.ID {
		t.Errorf("Expected query args to contain [%v, %v, %v, %v], instead got: args=%v", testUser.OrgID, testUser.Email, testUser.Password, testUser.ID, query.Args)
	}

}
