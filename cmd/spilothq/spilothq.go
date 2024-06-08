package main

import (
	"log"
	"net"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
	"github.com/aforamitdev/server-pilot/internal/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	informationService := &system.NewSystemInformer{}

	system_proto.RegisterInformerServer(grpcServer, informationService)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to server %s", err)
	}

}
