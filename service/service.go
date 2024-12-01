package service

import (
	"context"
	"errors"
	pb "github.com/dangrmous/transit-pdx-service/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

type Service struct {
	pb.UnimplementedTransitPDXServer
	serviceLogger Logger
}

func New(serviceLogger Logger) *Service {
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
