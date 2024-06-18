package collector

import (
	"fmt"
	"net"
)

type Collector struct {
	conn net.PacketConn
}

func NewCollector(address string) *Collector {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		fmt.Printf("failed to UDP listen on '%s' with '%v'", address, err)
	}

	return &Collector{conn: pc}
}

func (c *Collector) StreamTo(stream chan string) {
	buffer := make([]byte, 1024)

	go func() {
		n, addr, err := c.conn.ReadFrom(buffer)
		if err != nil {
			fmt.Printf("errore reading  address '%s'", addr)
		}
		fmt.Println(n)
		stream <- string(buffer)
	}()

}
