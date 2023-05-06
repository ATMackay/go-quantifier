package volatility

import (
	"github.com/ATMackay/go-quantifier/database"
	"github.com/ATMackay/go-quantifier/fetcher"
	"github.com/ATMackay/go-quantifier/rpc"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger  *logrus.Entry
	db      database.Database
	fetcher fetcher.Fetcher
	server  rpc.HTTPService
}

func (s *Service) Start() {
	s.logger.Info("startingVolatilityService")
	s.server.Start()
}

func (s *Service) Stop() {
	s.logger.Info("stoppingVolatilityService")
	if err := s.server.Stop(); err != nil {
		s.logger.WithFields(logrus.Fields{"error": err}).Error("errorStoppingServer")
	}
}
