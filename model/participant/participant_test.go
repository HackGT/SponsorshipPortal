package participant

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

func TestByRegistrationID(t *testing.T) {
	getReqChan := make(chan *modeltesting.QueryRequest, 1)
	getRespChan := make(chan error, 1)
	db := &modeltesting.MockConnection{
		GetReq:  getReqChan,
		GetResp: getRespChan,
	}

	// success case
	regID := "id1"
	getRespChan <- nil
	_, found, err := ByRegistrationID(db, regID)
	if err != nil || found != true {
		t.Errorf("Expected no error and found=true, instead got: found=%v, err=%v", found, err)
	}
	query := <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != regID {
		t.Errorf("Expected query args to contain only regID=%v, instead got: args=%v", regID, query.Args)
	}

	// error - no rows
	regID = "id2"
	getRespChan <- sql.ErrNoRows
	_, found, err = ByRegistrationID(db, regID)
	if err != sql.ErrNoRows || found != false {
		t.Errorf("Expected error=sql.ErrNoRows and found=false, instead got: found=%v, err=%v", found, err)
	}
	query = <-getReqChan
	if len(query.Args) != 1 || query.Args[0] != regID {
		t.Errorf("Expected query args to contain only regID=%v, instead got: args=%v", regID, query.Args)
	}
}
