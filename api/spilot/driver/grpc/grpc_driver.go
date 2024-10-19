package driver

import (
	"context"
	"fmt"
	"net"

	apiv1 "github.com/aforamitdev/server-pilot/app/spilothq/gen/proto/api/v1"
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

func (g *GrpcDriver) ConnectServer(ip net.IP, port string) (success bool, err error) {
	conn, err := grpc.NewClient(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println(conn.GetState())

	fmt.Println(err, "ERROR")
	g.conn = conn

	s := apiv1.NewSystemServicesClient(conn)

	req := apiv1.GetStatusRequest{}
	res, err := s.GetStatus(g.ctx, &req)
	if err != nil {
		fmt.Println(err, "Errr")
	}
	fmt.Println(res)

	return true, nil

}

func (g *GrpcDriver) GetServerStatus() apiv1.GetStatusResponse {

	ctx := context.Background()
	s := apiv1.NewSystemServicesClient(g.conn)

	req := apiv1.GetStatusRequest{}
	res, err := s.GetStatus(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}
	// if err != nil {
	// 	fmt.Println(err, "Errr")
	// }
	fmt.Println(res, "RES")
	return *res

}
