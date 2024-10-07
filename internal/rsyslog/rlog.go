package rsyslog

import (
	"fmt"
	"net"

	"github.com/aforamitdev/server-pilot/internal/protogen"
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

func (log *RLog) GetLogStream(*protogen.LogRequest, protogen.ServerPilot_GetLogStreamServer) error {
	fmt.Println("error ")
	return nil
}
