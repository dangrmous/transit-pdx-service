package service

import (
	"context"
	"errors"
	pb "github.com/dangrmous/transit-pdx-service/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Service struct {
	pb.UnimplementedTransitPDXServer
	serviceLogger *log.Logger
}

func New() *Service {
	return &Service{}
}

func (service *Service) Start(logger *log.Logger) error {
	service.serviceLogger = logger
	logger.Println("starting service")
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTransitPDXServer(grpcServer, New())
	grpcServer.Serve(lis)
	return errors.New("Service is not startable")
}

func (s *Service) GetScheduledTimes(context.Context, *pb.StopId) (*pb.NextScheduledTimes, error) {
	s.serviceLogger.Print("GetScheduledTimes called")
	times := []int32{123, 432}
	return &pb.NextScheduledTimes{
		ScheduledTimes: times,
	}, nil
}
