package rsyslog

import (
	"context"
	"fmt"
	"net"

	apiv1 "github.com/aforamitdev/server-pilot/app/spilothq/gen/proto/api/v1"
	"github.com/pkg/errors"
)

type RLog struct {
	// packet listener
	apiv1.UnimplementedLogServiceServer
	// protogen.ServerPilotServer

	PC net.PacketConn
}

func NewLogListener(port string) (*RLog, error) {

	pc, err := net.ListenPacket("udp", port)
	if err != nil {
		return nil, errors.Wrap(err, "fail to start UDP listeners, rsyslog")
	}
	return &RLog{PC: pc}, nil
}

func (log *RLog) GetLogs(req *apiv1.LogRequest, srv apiv1.LogService_GetLogsServer) error {

	// var rev protogen.LogRequest
	ctx := context.Background()

	data := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("CANCEL")
			return ctx.Err()
		default:
		}

		n, remoteAddr, err := log.PC.ReadFrom(data)
		if err != nil {
			fmt.Printf("error reading log")
		}
		prt := &apiv1.LogResponse{
			Log: string(data),
		}
		fmt.Println(remoteAddr)
		fmt.Println(n)

		err = srv.Send(prt)
		if err != nil {
			fmt.Println(err)
		}

	}

}
