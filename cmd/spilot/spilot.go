package main

import (
	"context"
	"log"
	"time"

	pb "github.com/aforamitdev/server-pilot/internal/proto/system"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// channel:=grpc.New

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error %s", err)

	}
	defer conn.Close()

	client := pb.NewInformerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetSystem(ctx, &pb.SystemRequest{})

	if err != nil {
		log.Fatalf("cound not greet : %v", err)
	}
	log.Printf("greeting: %s", r.GetName())

}
