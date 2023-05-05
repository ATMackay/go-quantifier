package rpc

import "net/http"

type Server interface {
	ListenAndServe() error
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}
