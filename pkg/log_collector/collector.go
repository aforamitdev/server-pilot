package log_collector

import (
	"errors"
	"fmt"
	"net"
)

type LogCollector struct {
	addr net.UDPAddr
	conn *net.UDPConn
}

func NewLogCollector() *LogCollector {
	return &LogCollector{}
}

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

	l.conn = conn
	return conn, nil
}

func (l *LogCollector) StreamLogs(logStream chan string) {
	// logStream := make(chan string)

	buf := make([]byte, 1024)

	rlan, remote, err := l.conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("error listening to log")
	}
	fmt.Println(remote)
	for {

		logStream <- string(buf[:rlan])
	}

}
