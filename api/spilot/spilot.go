package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aforamitdev/server-pilot/internal/protogen"
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
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(conn)
	if err != nil {
		log.Error(ctx, "startup", "msg", err)
	}

	defer conn.Close()

	c := protogen.NewServerPilotClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	r, err := c.GetLogStream(ctx, &protogen.LogRequest{})

	if err != nil {
		fmt.Println(err)
	}

	logchan := logstream(r)
	for res := range logchan {
		fmt.Println(res)
	}

	// r, err := client.

	// fmt.Println(r.Name)

	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	resp, err := stream.Recv()
	// 	if err == io.EOF {
	// 		return
	// 	} else if err == nil {
	// 		fmt.Printf(resp.String() + "\n")
	// 		// fmt.Println(valStr)
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}

func logstream(r protogen.ServerPilot_GetLogStreamClient) chan *protogen.LogResponse {
	logchan := make(chan *protogen.LogResponse)

	go func() {
		for {
			resp, err := r.Recv()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
			logchan <- resp
		}
	}()

	return logchan

}
