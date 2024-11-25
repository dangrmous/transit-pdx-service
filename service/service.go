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

func New(serviceLogger *log.Logger) *Service {
	return &Service{
		serviceLogger: serviceLogger,
	}
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
	pb.RegisterTransitPDXServer(grpcServer, service)
	grpcServer.Serve(lis)
	return errors.New("Service is not startable")
}

func (s *Service) GetScheduledTimes(c context.Context, sid *pb.StopId) (*pb.NextScheduledTimes, error) {
	s.serviceLogger.Printf("GetScheduledTimes sid=%s", sid)
	times := []int32{321}
	return &pb.NextScheduledTimes{
		ScheduledTimes: times,
	}, nil
}
