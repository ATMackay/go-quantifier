package rpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ATMackay/go-quantifier/logger"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HTTPService struct {
	port   int
	server Server
	logger *logrus.Entry
}

func NewHTTPService(port int, api *Api, l *logrus.Entry) HTTPService {
	handler := api.Routes()
	handler.Use(logger.LogHTTPRequest(l))

	return HTTPService{
		port: port,
		server: &http.Server{
			Addr:              fmt.Sprintf(":%d", port),
			Handler:           handler,
			ReadHeaderTimeout: 20 * time.Second,
		},
		logger: l,
	}
}

type EndPoint struct {
	path        string
	handlerFunc http.HandlerFunc
	methodType  string
}

func NewEndpoint(path, methodType string, handler http.HandlerFunc) EndPoint {
	return EndPoint{
		path:        path,
		handlerFunc: handler,
		methodType:  methodType,
	}
}

type Api struct {
	Endpoints []EndPoint
}

func (a *Api) AddEndpoint(e EndPoint) {
	a.Endpoints = append(a.Endpoints, e)
}

func (api Api) Routes() *mux.Router {
	router := mux.NewRouter()
	for _, e := range api.Endpoints {
		router.Handle(e.path, e.handlerFunc).Methods(e.methodType)
	}
	return router
}

func (service *HTTPService) Start() {
	go func() {
		if err := service.server.ListenAndServe(); err != nil {
			service.logger.WithFields(logrus.Fields{"error": err}).Warn("serverTerminated")
		}
	}()
}

func (service *HTTPService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return service.server.Shutdown(ctx)
}
