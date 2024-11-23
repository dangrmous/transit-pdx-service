package service

import (
	"errors"
	"log"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) Start(logger *log.Logger) error {
	logger.Println("starting service")
	return errors.New("Service is not startable")
}
