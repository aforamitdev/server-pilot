package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/aforamitdev/server-pilot/app/spilothq/services"
	"github.com/aforamitdev/server-pilot/internal/rsyslog"
	"github.com/aforamitdev/server-pilot/internal/system"
	"github.com/aforamitdev/server-pilot/pkg/logger"
	"github.com/aforamitdev/server-pilot/pkg/web"
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
	log.Info(ctx, "main: initializing API Support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// start rsys log listener
	rlog, err := rsyslog.NewLogListener(":5000")
	system := system.NewSystemInformer(log)
	if err != nil {
		log.Error(ctx, "error starting log listener")
		os.Exit(1)
	}

	server, err := services.NewServer(ctx, rlog, system)
	if err != nil {
		log.Error(ctx, "main:fail to init grpc server")
		os.Exit(1)

	}

	serviceErr := make(chan error, 1)

	go func() {
		serviceErr <- server.GrpcServer.Serve(lis)
	}()

	// err = server.Serve(lis)

	if err != nil {
		log.Info(ctx, "impossible to server %s", err)
	}

	select {
	case err := <-serviceErr:
		fmt.Println(err)
	case sig := <-shutdown:
		fmt.Println("closing service", sig)
		server.GrpcServer.Stop()
		os.Exit(1)

	}
}
