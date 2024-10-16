package services

import (
	"context"

	"github.com/aforamitdev/server-pilot/internal/rsyslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	// loge

	GrpcServer *grpc.Server
	rlog       *rsyslog.RLog
}

func NewServer(ctx context.Context, rlog *rsyslog.RLog) (*Server, error) {

	grpcServer := grpc.NewServer()
	grpcApiService := &Server{GrpcServer: grpcServer, rlog: rlog}

	reflection.Register(grpcServer)

	// protogen.RegisterServerPilotServer(grpcServer, rlog)

	return grpcApiService, nil

}
