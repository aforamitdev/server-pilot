package main

import (
	"context"
	"fmt"
	"net"
	"os"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
	"github.com/aforamitdev/server-pilot/internal/system"
	"github.com/aforamitdev/server-pilot/pkg/log_collector"
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

	// certificate file
	// cred, err := credentials.NewServerTLSFromFile("./cert/ca.cert", "./cert/ca.key")

	// if err != nil {
	// log.Error(ctx, "fail to generate cert file", err)
	// }

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	infoService := system.NewSystemInformer(log)
	system_proto.RegisterInformerServer(grpcServer, infoService)

	fmt.Printf("Server running...")
	logstream := make(chan string)

	lc := log_collector.NewLogCollector()
	conn, err := lc.StartLogCollector(5001, net.IP("0.0.0.0"))
	if err != nil {
		fmt.Println("error starting log collector")
	}
	fmt.Println(conn)
	lc.StreamLogs(logstream)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Info(ctx, "impossible to server %s", err)
	}

	select {
	case ls := <-logstream:
		fmt.Println(ls)

	}
}
