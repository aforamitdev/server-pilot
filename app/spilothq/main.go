package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/aforamitdev/server-pilot/app/spilothq/services"
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

	grpcService := services.Services("dev", shutdown, log)

	serviceErr := make(chan error, 1)

	go func() {
		serviceErr <- grpcService.Serve(lis)
	}()

	err = grpcService.Serve(lis)

	if err != nil {
		log.Info(ctx, "impossible to server %s", err)
	}

	select {
	case err := <-serviceErr:
		fmt.Println(err)
	case sig := <-shutdown:
		fmt.Println("closing service", sig)
		grpcService.Stop()
		os.Exit(1)

	}
}
