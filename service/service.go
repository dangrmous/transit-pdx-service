package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dangrmous/transit-pdx-service/config"
	pb "github.com/dangrmous/transit-pdx-service/pb"
	"github.com/dangrmous/transit-pdx-service/trimet"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Service struct {
	pb.UnimplementedTransitPDXServer
	serviceLogger Logger
}

// Returns a pointer to a new service
func New() *Service {
	return &Service{}
}

func (service *Service) Start(logger *log.Logger, conf *config.Config, tm trimet.TrimetClient) error {
	service.serviceLogger = logger
	logger.Println("starting service")
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", conf.Host, conf.Port))
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
