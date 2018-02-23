package testing

import (
	"database/sql"
	"errors"

	"github.com/HackGT/SponsorshipPortal/model"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrMockClosed     = errors.New("mock connection closed")
)

type QueryRequest struct {
	Dest  interface{}
	Query string
	Args  []interface{}
}

type SQLResultErrorPair struct {
	Result sql.Result
	Error  error
}

type MockConnection struct {
	ExecReq    chan *QueryRequest
	ExecResp   chan *SQLResultErrorPair
	GetReq     chan *QueryRequest
	GetResp    chan error
	SelectReq  chan *QueryRequest
	SelectResp chan error
}

var _ model.Connection = &MockConnection{}

func (m *MockConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.ExecReq == nil || m.ExecResp == nil {
		return nil, ErrNotImplemented
	}
	m.ExecReq <- &QueryRequest{Query: query, Args: args}
	result, ok := <-m.ExecResp
	if !ok {
		return nil, ErrMockClosed
	}
	return result.Result, result.Error
}

func (m *MockConnection) Get(dest interface{}, query string, args ...interface{}) error {
	if m.GetReq == nil || m.GetResp == nil {
		return ErrNotImplemented
	}
	m.GetReq <- &QueryRequest{dest, query, args}
	result, ok := <-m.GetResp
	if !ok {
		return ErrMockClosed
	}
	return result
}

func (m *MockConnection) Select(dest interface{}, query string, args ...interface{}) error {
	if m.SelectReq == nil || m.SelectResp == nil {
		return ErrNotImplemented
	}
	m.SelectReq <- &QueryRequest{dest, query, args}
	result, ok := <-m.SelectResp
	if !ok {
		return ErrMockClosed
	}
	return result
}
