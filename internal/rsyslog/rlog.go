package rsyslog

import (
	"net"

	"github.com/pkg/errors"
)

type RLog struct {
	// packet listener
	PC net.PacketConn
}

func NewLogListener(port string) (*RLog, error) {

	pc, err := net.ListenPacket("udp", port)
	if err != nil {
		return nil, errors.Wrap(err, "fail to start UDP listeners, rsyslog")
	}
	return &RLog{PC: pc}, nil
}
