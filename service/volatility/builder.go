package volatility

import (
	"net/http"

	"github.com/ATMackay/go-quantifier/fetcher"
	"github.com/sirupsen/logrus"
)

func BuildService(config Config) (*Service, error) {
	log := logrus.NewEntry(logrus.StandardLogger())
	h := http.Server{}
	rpcClient := http.Client{}
	f := fetcher.NewAlphaFetcher(&rpcClient)
	return &Service{logger: log, fetcher: f, server: &h}, nil
}
