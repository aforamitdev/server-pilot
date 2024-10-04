package services

import (
	"os"

	"github.com/aforamitdev/server-pilot/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Services(build string, shutdown chan os.Signal, log *logger.Logger) *grpc.Server {

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// infoService := system.NewSystemInformer(log)

	// system_proto.RegisterInformerServer(grpcServer, infoService)

	return grpcServer

}
