package services

import (
	"context"

	apiv1 "github.com/aforamitdev/server-pilot/app/spilothq/gen/proto/api/v1"
	"github.com/aforamitdev/server-pilot/internal/rsyslog"
	"github.com/aforamitdev/server-pilot/internal/system"
	"github.com/aforamitdev/server-pilot/pkg/logger"
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

func NewServer(ctx context.Context, log *logger.Logger) (*Server, error) {

	system := system.NewSystemInformer(log)
	rlog, err := rsyslog.NewLogListener(":5000")

	if err != nil {
		return &Server{}, err
	}

	grpcServer := grpc.NewServer()

	grpcApiService := &Server{GrpcServer: grpcServer, rlog: rlog, system: system}

	reflection.Register(grpcServer)

	apiv1.RegisterLogServiceServer(grpcServer, rlog)
	apiv1.RegisterSystemServicesServer(grpcServer, system)

	return grpcApiService, nil

}
