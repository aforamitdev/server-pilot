package rsyslog

import (
	"bytes"
	"fmt"
	"net"

	"github.com/pkg/errors"
)

type RLog struct {
	// packet listener
	protogen.UnimplementedServerPilotServer
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

func (log *RLog) GetLogStream(req *protogen.LogRequest, server protogen.ServerPilot_GetLogStreamServer) error {

	ctx := server.Context()

	// var rev protogen.LogRequest
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
		prt := protogen.LogResponse{
			Log: string(bytes.Trim(data, "\x00")),
		}
		fmt.Println(remoteAddr)
		fmt.Println(n)

		err = server.Send(&prt)
		if err != nil {
			fmt.Println(err)
		}

	}

}
