package driver

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcDriver struct {
	ctx  context.Context
	conn *grpc.ClientConn
}

func NewGrpcDriver() *GrpcDriver {
	return &GrpcDriver{}
}

func (g *GrpcDriver) Startup(ctx context.Context) {
	g.ctx = ctx
}

func (g *GrpcDriver) ConnectHQ(ip net.IP, port string) (success bool, err error) {
	conn, err := grpc.NewClient(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Println(fmt.Sprintf(`%s:%s`, ip, port))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(conn.GetState())
	fmt.Println(err, "ERROR")
	g.conn = conn

	return true, nil

}
