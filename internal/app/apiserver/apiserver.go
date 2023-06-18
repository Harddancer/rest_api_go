package apiserver

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Api server

type Apiserver struct {
	config *Config
	logger *logrus.Logger
}

// New
func New(config *Config) *Apiserver {
	fmt.Println("Done")
	return &Apiserver{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *Apiserver) Start() error {
	if err := s.cofigureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting API  server")
	return nil

}

func (s *Apiserver) cofigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
