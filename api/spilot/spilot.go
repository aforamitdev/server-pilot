package main

import (
	"context"
	"os"

	"github.com/aforamitdev/server-pilot/pkg/logger"
	"github.com/aforamitdev/server-pilot/pkg/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(ctx, "startup", "msg", err)
	}
	defer conn.Close()

}
