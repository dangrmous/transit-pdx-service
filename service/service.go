package service

import (
	"context"
	"errors"
	pb "github.com/dangrmous/transit-pdx-service/pb"
	"log"
)

type Service struct {
	pb.UnimplementedTransitPDXServer
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) Start(logger *log.Logger) error {
	logger.Println("starting service")
	return errors.New("Service is not startable")
}

func (s *Service) GetScheduledTimes(context.Context, *pb.StopId) (*pb.NextScheduledTimes, error) {
	times := []int32{123, 432}
	return &pb.NextScheduledTimes{
		ScheduledTimes: times,
	}, nil
}
