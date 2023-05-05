package volatility

import (
	"github.com/ATMackay/go-quantifier/fetcher"
	"github.com/ATMackay/go-quantifier/rpc"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Port      int    `yaml:"port"`
	Path      string `yaml:"path"`
	Loglevel  string `yaml:"loglevel"`
	LogFormat string `yaml:"logformat"`
}

type Service struct {
	logger  *logrus.Entry
	fetcher fetcher.Fetcher
	server  rpc.Server
}
