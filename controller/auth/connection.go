package auth

import (
	"database/sql"
	"errors"

	"github.com/HackGT/SponsorshipPortal/model"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrClosed = errors.New("connection closed")
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

type Connection struct {
	ExecReq    chan *QueryRequest
	ExecResp   chan *SQLResultErrorPair
	GetReq     chan *QueryRequest
	GetResp    chan error
	SelectReq  chan *QueryRequest
	SelectResp chan error
}

var _ model.Connection = &Connection{}

func (c *Connection) Exec(query string, args ...interface{}) (sql.Result, error) {
	if c.ExecReq == nil || c.ExecResp == nil {
		return nil, ErrNotImplemented
	}
	c.ExecReq <- &QueryRequest{Query: query, Args: args}
	result, ok := <-c.ExecResp
	if !ok {
		return nil, ErrClosed
	}
	return result.Result, result.Error
}

func (c *Connection) Get(dest interface{}, query string, args ...interface{}) error {
	if c.GetReq == nil || c.GetResp == nil {
		return ErrNotImplemented
	}
	c.GetReq <- &QueryRequest{dest, query, args}
	result, ok := <-c.GetResp
	if !ok {
		return ErrClosed
	}
	return result
}

func (c *Connection) Select(dest interface{}, query string, args ...interface{}) error {
	if c.SelectReq == nil || c.SelectResp == nil {
		return ErrNotImplemented
	}
	c.SelectReq <- &QueryRequest{dest, query, args}
	result, ok := <-c.SelectResp
	if !ok {
		return ErrClosed
	}
	return result
}
