package service

import (
	"log"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) Start(logger *log.Logger) error {
	logger.Println("starting service")
	return nil //errors.New("Service is not startable")
}
