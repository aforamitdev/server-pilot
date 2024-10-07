package services

import (
	"context"

	"github.com/aforamitdev/server-pilot/internal/rsyslog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	GrpcServer *grpc.Server
	rlog       *rsyslog.RLog
}

func NewServer(ctx context.Context, rlog *rsyslog.RLog) (*Server, error) {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	return &Server{GrpcServer: grpcServer, rlog: rlog}, nil

}
