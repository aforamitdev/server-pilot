package services

import (
	"context"
	"os"

	"github.com/aforamitdev/server-pilot/pkg/logger"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	echoServer *echo.Echo
}

func NewServer(ctx context.Context, grpcServer *grpc.Server) (*Server, error) {

	return &Server{grpcServer: grpcServer}, nil

}

func Services(build string, shutdown chan os.Signal, log *logger.Logger) *grpc.Server {

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// infoService := system.NewSystemInformer(log)

	// system_proto.RegisterInformerServer(grpcServer, infoService)

	return grpcServer

}
