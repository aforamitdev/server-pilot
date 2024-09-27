package log_collector

import (
	"errors"
	"net"
)

type LogCollector struct {
	addr net.UDPAddr
	conn *net.UDPConn
}

func NewLogCollector() {}

func (l *LogCollector) SetUpLog(port int, addressIP net.IP) {
	l.addr = net.UDPAddr{Port: port, IP: addressIP}
}

func (l *LogCollector) StartLogCollector(port int, address net.IP) (*net.UDPConn, error) {
	addr := net.UDPAddr{
		Port: port,
		IP:   address,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return nil, errors.New("fail to start log listener server")
	}
	defer conn.Close()
	return conn, nil
}

func (l *LogCollector) CloseLogCollector() {

}
