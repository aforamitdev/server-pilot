package main

import (
	"context"
	"net"
	"os"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
	"github.com/aforamitdev/server-pilot/internal/system"
	"github.com/aforamitdev/server-pilot/pkg/collector"
	"github.com/aforamitdev/server-pilot/pkg/logger"
	"github.com/aforamitdev/server-pilot/pkg/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	var log *logger.Logger
	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
		},
	}
	traceIDFn := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}
	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SPILOT-CLIENT", traceIDFn, events)
	ctx := context.Background()

	run(ctx, log)

}

func run(ctx context.Context, log *logger.Logger) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error(ctx, "cannot create listener: %s", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	infoService := system.NewSystemInformer(log)

	system_proto.RegisterInformerServer(grpcServer, infoService)

	err = grpcServer.Serve(lis)

	logChan := make(chan string, 2048)
	collector := collector.NewCollector("5000")
	collector.StreamTo(logChan)

	select {}

	if err != nil {
		log.Info(ctx, "impossible to server %s", err)
	}
}
