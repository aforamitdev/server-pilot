package system

import (
	"context"
	"fmt"
	"log"
	"net"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
	"github.com/aforamitdev/server-pilot/pkg/logger"
)

type SystemInformer struct {
	system_proto.UnimplementedInformerServer
	log *logger.Logger
}

func NewSystemInformer(log *logger.Logger) *SystemInformer {
	return &SystemInformer{
		log: log,
	}
}

func (s SystemInformer) GetSystem(ctx context.Context, req *system_proto.SystemRequest) (*system_proto.SystemResponse, error) {
	return &system_proto.SystemResponse{
		Name: "Amti rai",
	}, nil
}

func (s SystemInformer) GetLogs(req *system_proto.LogRequest, resp system_proto.Informer_GetLogsServer) error {

	addr := net.UDPAddr{
		Port: 5000,
		IP:   net.ParseIP("localhost"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {

		//time sleep to simulate server process time
		rlen, remote, err := conn.ReadFrom(buf)
		if err != nil {
			fmt.Println("error listening")
		}
		fmt.Println(rlen, remote)
		p := system_proto.LogResponse{Log: string(buf[:rlen])}

		if err := resp.Send(&p); err != nil {
			log.Printf("send error %v", err)
		}
	}

}
