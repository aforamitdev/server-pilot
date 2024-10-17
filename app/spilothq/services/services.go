package services

import (
	"context"

	apiv1 "github.com/aforamitdev/server-pilot/app/spilothq/gen/proto/api/v1"
	"github.com/aforamitdev/server-pilot/internal/rsyslog"
	"github.com/aforamitdev/server-pilot/internal/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	// loge
	apiv1.UnimplementedLogServiceServer
	apiv1.UnimplementedSystemServicesServer

	GrpcServer *grpc.Server
	rlog       *rsyslog.RLog
	system     *system.SystemInformer
}

func NewServer(ctx context.Context, rlog *rsyslog.RLog, system *system.SystemInformer) (*Server, error) {

	grpcServer := grpc.NewServer()
	grpcApiService := &Server{GrpcServer: grpcServer, rlog: rlog, system: system}

	reflection.Register(grpcServer)

	apiv1.RegisterLogServiceServer(grpcServer, grpcApiService)
	apiv1.RegisterSystemServicesServer(grpcServer, system)

	return grpcApiService, nil

}
