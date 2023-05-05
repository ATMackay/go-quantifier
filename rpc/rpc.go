package rpc

import (
	"context"
	"net/http"
)

type ResponseWriter http.ResponseWriter
type Request http.Request
type Response http.Response

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}
