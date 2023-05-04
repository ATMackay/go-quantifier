package rpc

type Server interface {
	ListenAndServe() error
}

type Client interface {
}
