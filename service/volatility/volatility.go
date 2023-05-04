package volatility

import (
	"net/http"

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
	logger *logrus.Entry
	server rpc.Server
}

func BuildService(config Config) (*Service, error) {
	log := logrus.NewEntry(logrus.StandardLogger())
	h := http.Server{}
	return &Service{logger: log, server: &h}, nil
}
