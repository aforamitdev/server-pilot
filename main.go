package main

import (
	"fmt"
	"net"
)

func main() {

	addr := net.UDPAddr{
		Port: 5000,
		IP:   net.ParseIP("localhost"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf [1024]byte

	for {
		rlen, remote, err := conn.ReadFrom(buf[:])
		if err != nil {
			fmt.Println("error listening")
		}
		fmt.Println(rlen)
		fmt.Println(remote)
		fmt.Println(string(buf[:]))

	}

}
