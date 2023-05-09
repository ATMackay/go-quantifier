package volatility

import (
	"net/http"

	"github.com/ATMackay/go-quantifier/database"
	"github.com/ATMackay/go-quantifier/fetcher"
	"github.com/ATMackay/go-quantifier/logger"
	"github.com/ATMackay/go-quantifier/rpc"
)

// BuildService creates the main service struct from config parameters
func BuildService(config Config) (*Service, error) {
	log, err := logger.NewLogger(logger.Level(config.Loglevel), logger.Format(config.LogFormat), false, ServiceName)
	if err != nil {
		return nil, err
	}
	db, err := database.NewBadgerWithLogger(config.Path, false, log.Logger)
	if err != nil {
		return nil, err
	}
	return &Service{
		logger:  log.Logger,
		db:      db,
		fetcher: fetcher.NewAlphaFetcher(&http.Client{}, config.ApiKey),
		server:  rpc.NewHTTPService(config.Port, makeAPIHandlers(), log.Logger)}, nil
}

func makeAPIHandlers() *rpc.Api {
	r := &rpc.Api{}
	r.AddEndpoint(rpc.NewEndpoint("/hello", http.MethodGet, Hello))
	return r
}
